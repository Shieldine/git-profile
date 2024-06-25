/*
Copyright © 2024 Shieldine <74987363+Shieldine@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"github.com/Shieldine/git-profile/internal"
	"github.com/spf13/cobra"
)

var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Reset credential config to none",
	Long: `Resets git attributes for current repository.
If you do this, git will default to your global config.`,
	Run: runUnset,
}

func runUnset(cmd *cobra.Command, args []string) {

	fmt.Println("warning: git will default to global credentials without local configuration")

	err := internal.UnsetUserName()
	if err != nil {
		fmt.Println(err)
	}

	err = internal.UnsetUserEmail()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Process complete.")
}

func init() {
	rootCmd.AddCommand(unsetCmd)
}
