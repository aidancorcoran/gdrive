/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

// */
// package cmd

// import (
// 	"fmt"

// 	"github.com/spf13/cobra"
// )

// // uploadCmd represents the upload command
// var uploadCmd = &cobra.Command{
// 	Use:   "upload",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("upload called")
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(uploadCmd)

// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// cmd/upload.go

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
