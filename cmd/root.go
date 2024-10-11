/*
Copyright © 2024 Aidan Corcoran <aidancorcoran.dev@gmail.com>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// root_cmd represents the base command when called without any subcommands
var root_cmd = &cobra.Command{
	Use:   "gdrive",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the root_cmd.
func Execute() {
	err := root_cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// root_cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gdrive.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	root_cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
