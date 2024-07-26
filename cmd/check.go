/*
Copyright © 2024 Shieldine <74987363+Shieldine@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"github.com/Shieldine/git-profile/internal"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Display the currently set credentials",
	Long:  `Check what credentials are currently set in the current project.`,
	Run:   runCheck,
}

func runCheck(cmd *cobra.Command, args []string) {
	name, err := internal.GetUserName()

	if err != nil {
		fmt.Println("error: not a git repository or username not set")
	} else {
		fmt.Printf("Current name: %s\n", name)
	}

	email, err := internal.GetUserEmail()

	if err != nil {
		fmt.Println("error: not a git repository or email not set")
	} else {
		fmt.Printf("Current email: %s\n", email)
	}
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
