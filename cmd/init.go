/*
Copyright © 2024 Shieldine <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Automatically set credentials for current repository",
	Long: `Use: git-profile init
to automatically set credentials for the current repository.
The credentials will be chosen by the repository's origin.

If no profile with a matching origin is present, you will be asked to 
add one with the possibility to use another profile as a template.

If multiple profiles with a matching origin are present, 
you will be asked to pick one.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
