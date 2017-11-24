// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package cmd

import (
	"fmt"

	"github.com/hoahm/mysh/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/mitchellh/mapstructure"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available servers",
	Long: `List all available servers.`,
	Aliases: []string{ "l" },
	Run: execList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func execList(cmd *cobra.Command, args []string) {
	fmt.Println("List all available servers")
	for prjName := range viper.GetStringMap("projects") {
		fmt.Println("----------------------------")
		fmt.Println("Project:", prjName)

		for id, m := range viper.GetStringMap("projects." + prjName) {
			var srv server.Server
			err := mapstructure.WeakDecode(m, &srv)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("\t - %s. %s\n", id, srv.Detail())
		}		
	}
	fmt.Println("")
}

// func prjList(m map[string]interface{}) []string {
// 	l := make([]string, 0)
// 	for prjName := range m {
// 		l = append(l, prjName)
// 	}
// 	return l
// }