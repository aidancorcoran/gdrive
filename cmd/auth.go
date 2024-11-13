/*
Copyright © 2024 Aidan Corcoran <aidancorcoran.dev@gmail.com>
*/

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Want to have users run gdrive auth in order to authenticate with their drive
// This command will have to be ran first in order for the other commands to work
var auth_cmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with your Google Drive Account",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("./token.json"); errors.Is(err, os.ErrNotExist) {
			fmt.Print("Token.json does not exist so execute auth command")
		} else {
			fmt.Print("Token.json does exist")
		}
	},
}

func init() {
	root_cmd.AddCommand(auth_cmd)
}
