/*
Copyright © 2024 Shieldine <74987363+Shieldine@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"github.com/Shieldine/git-profile/models"
	"os"

	"github.com/Shieldine/git-profile/internal"
	"github.com/spf13/cobra"
)

var all bool

var rmCmd = &cobra.Command{
	Use:   "rm [profile-name]",
	Short: "Remove existing profiles",
	Args:  cobra.MaximumNArgs(1),
	Long: `Provide <profile-name> to remove only the profile called <profile-name>.

Use --all flag to remove all profiles.
Use other flags to remove all profiles containing a specific name, email or origin.

<profile-name> and filtering flags cannot be provided together.

This action cannot be undone.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initialising profile removal...")

		if all {
			err := internal.ClearConfig()
			if err != nil {
				fmt.Println("Error removing all profiles:", err)
				os.Exit(1)
			}
			fmt.Println("All profiles removed from configuration.")
			return
		}

		if len(args) != 0 {
			if name != "" || email != "" || origin != "" {
				fmt.Println("Error: profile-name and flags cannot be provided together.")
				fmt.Println("Either provide a name or filtering options.")
				os.Exit(1)
			}

			profile := args[0]
			err := internal.DeleteProfile(profile)
			if err != nil {
				fmt.Printf("Error removing profile %s: %v\n", profile, err)
				os.Exit(1)
			}
			fmt.Printf("Profile %s removed.\n", profile)
			return
		}

		Profiles := internal.GetAllProfiles()
		var profiles []models.ProfileConfig

		profiles = append(profiles, Profiles...)

		if len(profiles) == 0 {
			fmt.Println("No profiles to remove.")
		} else {
			count := 0

			for _, profile := range profiles {
				if name != "" && name != profile.Name {
					continue
				} else if email != "" && email != profile.Email {
					continue
				} else if origin != "" && origin != profile.Origin {
					continue
				}

				err := internal.DeleteProfile(profile.ProfileName)
				if err != nil {
					fmt.Printf("Error removing profile %s: %v\n", profile.ProfileName, err)
					os.Exit(1)
				}

				fmt.Printf("Profile %s removed.\n", profile.ProfileName)
				count += 1
			}

			if count > 0 {
				fmt.Println()
				fmt.Printf("Successfuly removed %d profiles.\n", count)
			} else {
				fmt.Println("No profiles to remove.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
	rmCmd.Flags().BoolVarP(&all, "all", "a", false, "Remove all profiles")
	rmCmd.Flags().StringVarP(&name, "name", "n", "", "Remove profiles with name")
	rmCmd.Flags().StringVarP(&email, "email", "e", "", "Remove profiles with email")
	rmCmd.Flags().StringVarP(&origin, "origin", "o", "", "Remove profiles with origin")
}
