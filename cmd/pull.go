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

// Main command that handles all the logic with downloading Google Drive files
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

		file_id, mime_type := getFileIdAndMimeType(srv, file_name)

		file_extension, err := getFileExtension(mime_type)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		resp, err := downloadGoogleDriveFiles(srv, file_id, file_extension, file_name)
		if err != nil {
			log.Fatalf("Error: Downloading Google Drive Files: %v\n", err)
		}

		fmt.Printf("File '%s' exported and downloaded as '%s' successfully\n", file_name, resp)
	},
}

func downloadGoogleDriveFiles(srv *drive.Service, file_id string, file_extension string, file_name string) (string, error) {
	resp, err := srv.Files.Export(file_id, getExportMimeType(file_extension)).Download()
	if err != nil {
		return "", fmt.Errorf("unable to export file: %v", err)
	}
	defer resp.Body.Close()

	// Need to handle weird file names
	local_file_name := file_name + file_extension
	out_file, err := os.Create(local_file_name)
	if err != nil {
		return "", fmt.Errorf("unable to create local file: %v", err)
	}
	defer out_file.Close()

	_, err = io.Copy(out_file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to save local file: %v", err)
	}

	return local_file_name, nil
}

// We error checked the file extension in getFileExtension()
func getExportMimeType(file_extension string) string {
	// Determine the MIME type from the extension
	for key, val := range other_mime_types {
		if val == file_extension {
			return key
		}
	}
	return ""
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
	return file.Id, file.MimeType
}

func getFileExtension(mime_type string) (string, error) {
	if extension, exists := gdrive_mime_types[mime_type]; exists {
		return extension, nil
	} else if extension, exists := other_mime_types[mime_type]; exists {
		return extension, nil
	}
	return "", fmt.Errorf("unknown MIME Type: %s", mime_type)
}

func init() {
	root_cmd.AddCommand(pull_cmd)
}
