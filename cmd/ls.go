/*
Copyright © 2024 Shieldine <74987363+Shieldine@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"github.com/Shieldine/git-profile/internal"
	"github.com/Shieldine/git-profile/models"
	"github.com/spf13/cobra"
)

var (
	profileName = ""
	name        string
	email       string
	origin      string
)

var lsCmd = &cobra.Command{
	Use:     "list [profile-name]",
	Aliases: []string{"l", "ls"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "List profiles",
	Long: `Display profiles currently present in your config.

Provide a profile name to list the attributes of the specified profile.
Use flags to filter for a specific origin, name or email.
`,
	Run: runLs,
}

func runLs(_ *cobra.Command, args []string) {
	if len(args) != 0 {
		profileName = args[0]

		Profile := internal.GetProfileByName(profileName)

		if (models.ProfileConfig{}) == Profile {
			fmt.Printf("Profile %s doesn't exist.", profileName)
			return
		}

		PrintProfile(Profile)
		return
	}

	Profiles := internal.GetAllProfiles()

	if len(Profiles) == 0 {
		fmt.Println("No profiles to display.")
	} else {
		for _, profile := range Profiles {
			if name == "" && email == "" && origin == "" {
				PrintProfile(profile)
				continue
			} else if name != "" && name != profile.Name {
				continue
			} else if email != "" && email != profile.Email {
				continue
			} else if origin != "" && origin != profile.Origin {
				continue
			}
			PrintProfile(profile)
		}
	}
}

func PrintProfile(profile models.ProfileConfig) {
	fmt.Printf("Profile %s:\n", profile.ProfileName)
	fmt.Printf("  Origin: %s\n", profile.Origin)
	fmt.Printf("  Name: %s\n", profile.Name)
	fmt.Printf("  Email: %s\n", profile.Email)
	fmt.Println()
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().StringVarP(&name, "name", "n", "", "List profiles with matching name")
	lsCmd.Flags().StringVarP(&email, "email", "e", "", "List profiles with matching email")
	lsCmd.Flags().StringVarP(&origin, "origin", "o", "", "List profiles with matching origin")
}
