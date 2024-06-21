/*
Copyright © 2024 Shieldine <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Shieldine/git-profile/internal"
	"github.com/Shieldine/git-profile/models"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:     "set <profile-name>",
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(1),
	Short:   "Set profile for current repository",
	Long: `Change the current repository's profile to <profile-name>.

If the origin is different than what is set in the profile, 
you will be asked if you want to create a new profile for this origin.
`,
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]
		profile := internal.GetProfileByName(profileName)

		if (models.ProfileConfig{}) == profile {
			reader := bufio.NewReader(os.Stdin)

			fmt.Printf("Profile %s doesn't exist.\n", profileName)
			fmt.Print("Would you like to create it? (y/n): ")

			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(answer)

			if answer == "n" {
				fmt.Println("Nothing to do.")
				return
			} else if answer == "y" {
				addRun(cmd, []string{profileName})
			} else {
				fmt.Println("Invalid choice.")
				os.Exit(1)
			}
		}

		profile = internal.GetProfileByName(profileName)
		currentName, err := internal.GetUserName()
		currentEmail, _ := internal.GetUserEmail()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if profile.Name == currentName && profile.Email == currentEmail {
			fmt.Println("Nothing to do.")
			return
		}

		err = internal.SetUserName(profile.Name)
		if err != nil {
			fmt.Printf("Error setting user name: %s\n", err)
			os.Exit(1)
		}

		err = internal.SetUserEmail(profile.Email)
		if err != nil {
			fmt.Printf("Error setting user email: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Profile %s set.\n", profileName)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
