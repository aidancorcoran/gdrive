/*
Copyright © 2024 Aidan Corcoran <aidancorcoran.dev@gmail.com>
*/

package cmd

import (
	"fmt"
	"log"

	"gdrive/pkg/auth"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var flag_queries = map[string]string{
	"shared": "sharedWithMe",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List files in Google Drive",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := auth.GetDriveService()
		if err != nil {
			log.Fatalf("Unable to retrieve Drive client: %v", err)
		}

		var query string
		for key, value := range flag_queries {
			flag_exist, err := cmd.Flags().GetBool(key)
			if err != nil {
				log.Fatalf("Unable to retrieve files: %v", err)
			} else if flag_exist {
				query += value
			}
		}

		//"'me' in owners"
		r, err := srv.Files.List().Q(query).Do()
		// .PageSize(50).
		// 	Fields("nextPageToken, files(id, name)").Do()
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

func handleFlags(flags *pflag.FlagSet) (query string) {
	// var query string

	return query
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("shared", "s", false, "List only files shared with your Google account.")
}
