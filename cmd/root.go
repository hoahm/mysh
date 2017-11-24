// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package cmd

import (
	"fmt"
	"os"

	"github.com/hoahm/mysh/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CfgFileName default config file name
const CfgFileName = ".mysh"

var (
	cfgFile string
	// Verbose the message
	Verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mysh",
	Short: "A simple SSH client help you quickly make a SSH connection to your servers.",
	Long: `A simple SSH client help you quickly make a SSH connection to your servers, without having to remember the server names, IPs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() { 
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mysh.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "help message for toggle")

	// Enable suggestion
	rootCmd.DisableSuggestions = false
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else if cfgFileExists() {
		// Use config file from current directory
		curDir := util.CurrentDir()
		viper.AddConfigPath(curDir)
	} else {
		// Use config file from home directory.
		home := util.HomeDir()
		viper.AddConfigPath(home)
	}

	// Search config in home directory with name ".mysh" (without extension)
	viper.SetConfigName(CfgFileName)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file %s", viper.ConfigFileUsed())
	}
}

func cfgFileExists() bool {
	curDir := util.CurrentDir()
	ymlFile := util.FilePath(curDir, CfgFileName, "yml")
	yamlFile := util.FilePath(curDir, CfgFileName, "yaml")

	if util.Exists(ymlFile) || util.Exists(yamlFile) {
		return true
	}
	return false
}