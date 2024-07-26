/*
Copyright © 2024 Shieldine <74987363+Shieldine@users.noreply.github.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/Shieldine/git-profile/internal"
	"github.com/Shieldine/git-profile/models"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var editCmd = &cobra.Command{
	Use:     "update <profile-name>",
	Aliases: []string{"edit", "u", "e"},
	Args:    cobra.ExactArgs(1),
	Short:   "Update an existing profile",
	Long: `Update the profile <profile-name>.
You will be asked to update your credentials and origin.
The updated values can be passed as flags.
`,
	Run: runUpdate,
}

func runUpdate(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)
	profileName := args[0]

	oldProfile := internal.GetProfileByName(profileName)

	if (models.ProfileConfig{}) == oldProfile {
		fmt.Printf("Profile %s doesn't exist.\n", profileName)
		return
	}

	if name == "" {

		fmt.Printf("Name (enter to keep %s): ", oldProfile.Name)
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			name = oldProfile.Name
		}
	}

	if email == "" {
		fmt.Printf("E-mail (enter to keep %s): ", oldProfile.Email)
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)

		if email == "" {
			email = oldProfile.Email
		}
	}

	currentOrigin, _ := internal.GetRepoOrigin()
	newOrigin := ""

	if origin == "" {
		fmt.Printf("Origin (enter to keep %s): ", oldProfile.Origin)

		newOrigin, _ = reader.ReadString('\n')
		newOrigin = strings.TrimSpace(newOrigin)

		if newOrigin == "" {
			newOrigin = oldProfile.Origin
		}
	} else {
		if origin == "auto" {
			newOrigin = currentOrigin
		} else {
			newOrigin = origin
		}
	}

	err := internal.EditProfile(args[0], models.ProfileConfig{
		ProfileName: profileName,
		Name:        name,
		Email:       email,
		Origin:      newOrigin,
	})
	if err != nil {
		fmt.Printf("Error updating profile: %v\n", err)
		return
	}

	fmt.Printf("Profile %s updated\n", profileName)
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringVarP(&name, "name", "n", "", "Set the name directly")
	editCmd.Flags().StringVarP(&email, "email", "e", "", "Set the email directly")
	editCmd.Flags().StringVarP(&origin, "origin", "o", "", "Set the origin directly."+
		" Type \"auto\" to accept origin of the current repository")
}
