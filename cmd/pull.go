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
