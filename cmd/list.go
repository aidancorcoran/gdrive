/*
// Copyright © 2024 NAME HERE <EMAIL ADDRESS>

// */
// package cmd

// import (
// 	"fmt"

// 	"github.com/spf13/cobra"
// )

// // listCmd represents the list command
// var listCmd = &cobra.Command{
// 	Use:   "list",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("list called")
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(listCmd)

// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// cmd/list.go

package cmd

import (
	"fmt"
	"log"

	"gdrive/pkg/auth"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List files in Google Drive",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := auth.GetDriveService()
		if err != nil {
			log.Fatalf("Unable to retrieve Drive client: %v", err)
		}

		r, err := srv.Files.List().PageSize(10).
			Fields("nextPageToken, files(id, name)").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve files: %v", err)
		}

		fmt.Println("Files:")
		if len(r.Files) == 0 {
			fmt.Println("No files found.")
		} else {
			for _, i := range r.Files {
				fmt.Printf("%s (%s)\n", i.Name, i.Id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
