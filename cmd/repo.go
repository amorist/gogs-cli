package cmd

import (
	"fmt"
	"gogs-cli/gogs"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	repoCmd.AddCommand(listReposCmd)
	listReposCmd.Flags().StringP("name", "n", "", "User name")
}

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage repos",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var listReposCmd = &cobra.Command{
	Use:   "list",
	Short: "list repositories",
	Run: func(cmd *cobra.Command, args []string) {
		var repositorys []*gogs.Repository
		var err error
		var name string

		name, err = cmd.Flags().GetString("name")

		if name == "" {
			repositorys, err = client.ListMyRepos()
		} else {
			repositorys, err = client.ListUserRepos(name)
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Description", "Owner"})

		for _, repository := range repositorys {
			info := []string{}
			info = append(info, repository.Name)
			info = append(info, repository.Description)
			info = append(info, repository.Owner.UserName)
			table.Append(info)
		}
		table.Render()
	},
}
