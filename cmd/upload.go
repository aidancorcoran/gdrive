/*
Copyright © 2024 Aidan Corcoran <aidancorcoran.dev@gmail.com>
*/

package cmd

import (
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"

	"gdrive/pkg/auth"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

var upload_cmd = &cobra.Command{
	Use:   "upload [file path]",
	Short: "Upload a file to Google Drive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file_path := args[0]
		srv, err := auth.GetDriveService()
		if err != nil {
			log.Fatalf("Unable to retrieve Drive client: %v", err)
		}

		f, err := os.Open(file_path)
		if err != nil {
			log.Fatalf("Error opening %q: %v", file_path, err)
		}
		defer f.Close()

		file_name := filepath.Base(file_path)
		mime_type := mime.TypeByExtension(filepath.Ext(file_name))
		if mime_type == "" {
			mime_type = "application/octet-stream"
		}

		file := &drive.File{Name: file_name}
		res, err := srv.Files.Create(file).Media(f, googleapi.ContentType(mime_type)).Do()
		if err != nil {
			log.Fatalf("Error uploading file: %v", err)
		}

		fmt.Printf("File '%s' uploaded successfully. ID: %s\n", res.Name, res.Id)
	},
}

func init() {
	root_cmd.AddCommand(upload_cmd)
}
