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

var uploadCmd = &cobra.Command{
	Use:   "upload [file path]",
	Short: "Upload a file to Google Drive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		srv, err := auth.GetDriveService()
		if err != nil {
			log.Fatalf("Unable to retrieve Drive client: %v", err)
		}

		f, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Error opening %q: %v", filePath, err)
		}
		defer f.Close()

		fileName := filepath.Base(filePath)
		mimeType := mime.TypeByExtension(filepath.Ext(fileName))
		if mimeType == "" {
			mimeType = "application/octet-stream"
		}

		file := &drive.File{Name: fileName}
		res, err := srv.Files.Create(file).Media(f, googleapi.ContentType(mimeType)).Do()
		if err != nil {
			log.Fatalf("Error uploading file: %v", err)
		}

		fmt.Printf("File '%s' uploaded successfully. ID: %s\n", res.Name, res.Id)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
