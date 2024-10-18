/*
Copyright © 2024 Aidan Corcoran <aidancorcoran.dev@gmail.com>
*/

package cmd

import (
	"fmt"
	"log"

	"github.com/aidancorcoran/gdrive/pkg/auth"

	"github.com/spf13/cobra"
)

var flag_queries = map[string]string{
	"shared": "sharedWithMe",
}

var list_cmd = &cobra.Command{
	Use:   "list",
	Short: "List files in Google Drive",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := auth.GetDriveService()
		if err != nil {
			log.Fatalf("Unable to retrieve Drive client: %v", err)
		}

		var query string = handleFlags(cmd)
		file_list, err := srv.Files.List().Q(query).Do()
		// .PageSize(50).
		// 	Fields("nextPageToken, files(id, name)").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve files: %v", err)
		}

		fmt.Println("Files:")
		if len(file_list.Files) == 0 {
			fmt.Println("No files found.")
		} else {
			for _, i := range file_list.Files {
				fmt.Printf("%s (%s)\n", i.Name, i.Id)
			}
		}
	},
}

func handleFlags(cmd *cobra.Command) (query string) {
	// Iterate over the possible list command flags and determine if they are present
	for key, value := range flag_queries {
		flag_exist, err := cmd.Flags().GetBool(key)
		if err != nil {
			log.Fatalf("Unable to retrieve flags: %v", err)
		} else if flag_exist { // If the flag exists append its query value
			query += value
		} else {
			// Default query command when just running gdrive list
			query += "'me' in owners and trashed=false"
		}
	}

	return query
}

func init() {
	root_cmd.AddCommand(list_cmd)

	list_cmd.Flags().BoolP("shared", "s", false, "List only files shared with your Google account.")
}
