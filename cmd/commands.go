package cmd

import (
	"fmt"
	"gogs-cli/gogs"
	"os"

	"github.com/spf13/cobra"
)

var client *gogs.Client

func init() {
	GOGSURL := os.Getenv("GOGS_SERVER")
	GOGSTOKEN := os.Getenv("GOGS_TOKEN")
	client = gogs.NewClient(GOGSURL, GOGSTOKEN)
	rootCmd.AddCommand(userCmd, repoCmd, versionCmd)
}

func version() string {
	return "0.0.1"
}

var rootCmd = &cobra.Command{
	Use:   "gogs",
	Short: "gogs cli",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gogs cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version())
	},
}

// Execute cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
