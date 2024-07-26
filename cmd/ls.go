/*
Copyright © 2024 Shieldine <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Shieldine/git-profile/internal"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "List profiles",
	Long:    `Display all profiles currently present in your config.`,
	Run: func(cmd *cobra.Command, args []string) {
		Users := internal.GetAllProfiles()
		if len(Users) == 0 {
			fmt.Println("No profiles to display.")
		} else {
			for _, user := range Users {
				fmt.Println(user)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
