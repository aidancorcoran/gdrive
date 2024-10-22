/*
Copyright © 2024 Aidan Corcoran <aidancorcoran.dev@gmail.com>
*/

package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aidancorcoran/gdrive/pkg/auth"
	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// Convert Google Drive Specific MIME types to file extensions
var gdrive_mime_types = map[string]string{
	"application/vnd.google-apps.audio":        ".mp3",
	"application/vnd.google-apps.document":     ".docx",
	"application/vnd.google-apps.drive-sdk":    ".unknown",
	"application/vnd.google-apps.drawing":      ".png",
	"application/vnd.google-apps.file":         ".unknown",
	"application/vnd.google-apps.folder":       ".folder",
	"application/vnd.google-apps.form":         ".form",
	"application/vnd.google-apps.fusiontable":  ".table",
	"application/vnd.google-apps.jam":          ".jam",
	"application/vnd.google-apps.mail-layout":  ".email",
	"application/vnd.google-apps.map":          ".map",
	"application/vnd.google-apps.photo":        ".jpg",
	"application/vnd.google-apps.presentation": ".pptx",
	"application/vnd.google-apps.script":       ".js",
	"application/vnd.google-apps.shortcut":     ".shortcut",
	"application/vnd.google-apps.site":         ".html",
	"application/vnd.google-apps.spreadsheet":  ".xlsx",
	"application/vnd.google-apps.unknown":      ".unknown",
	"application/vnd.google-apps.vid":          ".mp4",
	"application/vnd.google-apps.video":        ".mp4",
}

// Convert other MIME types to file extensions
var other_mime_types = map[string]string{
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": ".docx",
	"application/vnd.oasis.opendocument.text":                                 ".odt",
	"application/rtf":      ".rtf",
	"application/pdf":      ".pdf",
	"text/plain":           ".txt",
	"application/zip":      ".zip",
	"application/epub+zip": ".epub",
	"text/markdown":        ".md",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": ".xlsx",
	"application/x-vnd.oasis.opendocument.spreadsheet":                  ".ods",
	"text/csv":                  ".csv",
	"text/tab-separated-values": ".tsv",
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": ".pptx",
	"application/vnd.oasis.opendocument.presentation":                           ".odp",
	"image/jpeg":    ".jpg",
	"image/png":     ".png",
	"image/svg+xml": ".svg",
	"application/vnd.google-apps.script+json": ".json",
}

var pull_cmd = &cobra.Command{
	Use:   "pull [file name]",
	Short: "Download a file from Google Drive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file_name := args[0]
		srv, err := auth.GetDriveService()
		if err != nil {
			log.Fatalf("Unable to retrieve Drive client: %v", err)
		}

		// Need to work on converting the mime_type from the Google Api to an appropriate exportable value
		// https://developers.google.com/drive/api/guides/ref-export-formats
		file_id, mime_type := getFileIdAndMimeType(srv, file_name)

		// Convert this functionality to a function
		if value, exists := gdrive_mime_types[mime_type]; exists {
			fmt.Printf("GDrive Specific Mime Type: %s\nValue: %s\n", mime_type, value)
		} else if value, exists := other_mime_types[mime_type]; exists {
			fmt.Printf("Other mime type: %s\nValue: %s\n", mime_type, value)
		} else {
			fmt.Printf("Mime type %s does not exist\n", mime_type)
		}

		// Add logic to handle export the file if it is a gdrive_mime_type and a get function if it is a other mime type
		resp, err := srv.Files.Export(file_id, "application/pdf").Download()
		if err != nil {
			log.Fatalf("Unable to export file: %v", err)
		}
		defer resp.Body.Close()

		// Create a local file to save the exported content
		localFileName := file_name + ".pdf"
		outFile, err := os.Create(localFileName)
		if err != nil {
			log.Fatalf("Unable to create local file: %v", err)
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, resp.Body)
		if err != nil {
			log.Fatalf("Unable to save file: %v", err)
		}

		fmt.Printf("File '%s' exported and downloaded as '%s' successfully\n", file_name, localFileName)
	},
}

func getFileIdAndMimeType(srv *drive.Service, file_name string) (string, string) {
	query := fmt.Sprintf("name = '%s'", file_name)

	file_list, err := srv.Files.List().Q(query).Fields("files(id, mimeType)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	if len(file_list.Files) == 0 {
		return "", ""
	}

	file := file_list.Files[0]
	fmt.Printf("Found file: %s (ID: %s, MIME Type: %s)\n", file_name, file.Id, file.MimeType)
	return file.Id, file.MimeType
}

func init() {
	root_cmd.AddCommand(pull_cmd)
}
