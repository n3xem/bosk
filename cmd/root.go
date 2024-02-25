/*
Copyright © 2024 yukyan
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bosk",
	Short: "bosk is a tool to generate ssh key pairs and store them in a file.",
	Long:  `bosk is a tool to generate ssh key pairs and store them in a file. 'bosk' is abbreviation of "Bunch Of SSH Key pairs".`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bosk.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
