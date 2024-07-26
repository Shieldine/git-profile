/*
Copyright © 2024 Shieldine <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Shieldine/git-profile/internal"
	"os"

	"github.com/spf13/cobra"
)

var all bool

var rmCmd = &cobra.Command{
	Use:   "rm <profile-name> | rm --all",
	Short: "Remove existing profiles",
	Args:  cobra.MaximumNArgs(1),
	Long: `Remove the profile called <profile-name>.
Use --all flag to remove all profiles.
This action cannot be undone.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if all {
			err := internal.ClearConfig()
			if err != nil {
				fmt.Println("Error removing all users:", err)
				os.Exit(1)
			}
			fmt.Println("All profiles removed from configuration.")
		} else {
			if len(args) == 0 {
				fmt.Println("Error: Requires a username or --all flag.")
				os.Exit(1)
			}
			profile := args[0]
			err := internal.DeleteProfile(profile)
			if err != nil {
				fmt.Printf("Error removing profile %s: %v\n", profile, err)
				os.Exit(1)
			}
			fmt.Printf("Profile %s removed.\n", profile)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
	rmCmd.Flags().BoolVarP(&all, "all", "a", false, "Remove all profiles")
}
