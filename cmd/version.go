// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VERSION Hold the current version of this application
const VERSION = "0.0.1"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long: `Print the version.`,
	Aliases: []string{ "v" },
	Example: `
	mysh version
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
