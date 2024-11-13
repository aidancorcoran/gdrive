/*
Copyright © 2024 Aidan Corcoran <aidancorcoran.dev@gmail.com>
*/

package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aidancorcoran/gdrive/pkg/auth"
	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

var drive_service *drive.Service

// Want to have users run gdrive auth in order to authenticate with their drive
// This command will have to be ran first in order for the other commands to work
var auth_cmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with your Google Drive Account",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("./token.json"); errors.Is(err, os.ErrNotExist) {
			drive_service, err := auth.GetDriveService()
			if err != nil {
				log.Fatalf("Unable to retrieve Drive client: %v", err)
			}
			fmt.Println("Authentication successful. Drive service initialized.")
		} else {
			fmt.Print("Token.json does exist")
		}
	},
}

func ActiveAccount() *drive.Service {
	return drive_service
}

func init() {
	root_cmd.AddCommand(auth_cmd)
}
