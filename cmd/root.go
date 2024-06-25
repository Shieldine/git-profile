/*
Copyright © 2024 Shieldine <74987363+Shieldine@users.noreply.github.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-profile",
	Short: "Manage and automatically set git user profiles based on the project's origin",
	Long: `git-profile is a simple CLI to manage and automatically set git user profiles based on the project's origin.

Save a profile together with its origin and let git-profile set the attributes next time you clone a new repository.
To make managing names and emails more convenient in general, git-profile offers further commands that will let you
check, unset and set credentials without creating a profile.
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
