package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	userCmd.AddCommand(userInfoCmd)
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var userInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get user info",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			user, err := client.GetUserInfo(args[0])
			if err != nil {
				fmt.Println(err)
			}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Name", "Email"})
			info := []string{user.UserName, user.Email}
			table.Append(info)
			table.Render()
		}
	},
}
