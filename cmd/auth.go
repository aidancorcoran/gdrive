/*
Copyright © 2024 Aidan Corcoran <aidancorcoran.dev@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// Want to have users run gdrive auth in order to authenticate with their drive
// This command will have to be ran first in order for the other commands to work
var auth_cmd = &cobra.Command{}
