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

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Automatically set credentials for current repository",
	Long: `Automatically set credentials for the current repository.
The credentials will be chosen by the repository's origin.

If no profile with a matching origin is present, you will be asked to 
add one.

If multiple profiles with a matching origin are present, 
you will be asked to pick one.
`,
	Run: runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) {
	currentOrigin, err := internal.GetRepoOrigin()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	possibleProfiles := internal.GetProfilesByOrigin(currentOrigin)

	if len(possibleProfiles) == 0 {
		fmt.Printf("No profiles found for origin %s\n", currentOrigin)
		fmt.Print("Would you like to create a new one? (y/n): ")

		answer := ReadAnswer()

		if answer == "n" {
			fmt.Println("Nothing to do")
			return
		} else {
			addRun(cmd, []string{})
		}

		possibleProfiles = internal.GetProfilesByOrigin(currentOrigin)

		if CredentialsAlreadySet(possibleProfiles[0]) {
			fmt.Println("Repository already has correct credentials. Nothing to do.")
			return
		}

		err = internal.SetUserName(possibleProfiles[0].Name)
		if err != nil {
			fmt.Println(err)
		}

		err = internal.SetUserEmail(possibleProfiles[0].Email)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Credentials of profile %s set for current project.\n", possibleProfiles[0].ProfileName)

	} else if len(possibleProfiles) == 1 {
		if CredentialsAlreadySet(possibleProfiles[0]) {
			fmt.Println("Repository already has correct credentials. Nothing to do.")
			return
		}

		err = internal.SetUserName(possibleProfiles[0].Name)
		if err != nil {
			fmt.Println(err)
		}

		err = internal.SetUserEmail(possibleProfiles[0].Email)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Credentials of profile %s set for current project.\n", possibleProfiles[0].ProfileName)
	} else {
		fmt.Printf("Multiple profiles found for origin %s\n", currentOrigin)
		for _, possibleProfile := range possibleProfiles {
			PrintProfile(possibleProfile)
		}
		fmt.Println("Please pick a profile (enter the profile name):")

		reader := bufio.NewReader(os.Stdin)

		selectedProfile := models.ProfileConfig{}

		for {
			profileName, _ = reader.ReadString('\n')
			profileName = strings.TrimSpace(profileName)

			fits := false
			for _, possibleProfile := range possibleProfiles {
				if profileName == possibleProfile.ProfileName {
					fits = true
					selectedProfile = possibleProfile
				}
			}
			if fits {
				break
			} else {
				fmt.Println("Invalid choice. Please try again.")
			}
		}

		if CredentialsAlreadySet(selectedProfile) {
			fmt.Println("Repository already has correct credentials. Nothing to do.")
			return
		}

		err = internal.SetUserName(selectedProfile.Name)
		if err != nil {
			fmt.Println(err)
		}

		err = internal.SetUserEmail(selectedProfile.Email)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Credentials of profile %s set for current project.\n", selectedProfile.ProfileName)
	}
}

func CredentialsAlreadySet(profile models.ProfileConfig) bool {
	currentName, _ := internal.GetUserName()
	currentEmail, _ := internal.GetUserEmail()

	if profile.Name == currentName && profile.Email == currentEmail {
		return true
	}
	return false
}
