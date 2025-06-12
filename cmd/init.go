// Package cmd
// Copyright © 2024 Shieldine
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// /*
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

// initCmd represents the init command for automatically setting git attributes
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Automatically set attributes for current repository",
	Long: `Automatically set attributes for the current repository.
The attributes will be chosen by the repository's origin.

If no profile with a matching origin is present, you will be asked to 
add one.

If multiple profiles with a matching origin are present, 
you will be asked to pick one.

Usage:
  git-profile init
`,
	Run: runInit,
}

// runInit handles the init command execution.
// It automatically sets Git credentials based on the repository's origin.
// The function follows these steps:
// 1. Get the current repository's origin
// 2. Find profiles matching that origin
// 3. If no matching profiles, prompt to create one
// 4. If one matching profile, use it
// 5. If multiple matching profiles, ask user to select one
func runInit(cmd *cobra.Command, _ []string) {
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
			runAdd(cmd, []string{})
		}

		possibleProfiles = internal.GetProfilesByOrigin(currentOrigin)

		if CredentialsAlreadySet(possibleProfiles[0]) {
			fmt.Println("Repository already has correct credentials. Nothing to do.")
			return
		}

		err = internal.SetUserName(possibleProfiles[0].Name, false)
		if err != nil {
			fmt.Println(err)
		}

		err = internal.SetUserEmail(possibleProfiles[0].Email, false)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Credentials of profile %s set for current project.\n", possibleProfiles[0].ProfileName)

	} else if len(possibleProfiles) == 1 {
		if CredentialsAlreadySet(possibleProfiles[0]) {
			fmt.Println("Repository already has correct credentials. Nothing to do.")
			return
		}

		err = internal.SetUserName(possibleProfiles[0].Name, false)
		if err != nil {
			fmt.Println(err)
		}

		err = internal.SetUserEmail(possibleProfiles[0].Email, false)
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

		err = internal.SetUserName(selectedProfile.Name, false)
		if err != nil {
			fmt.Println(err)
		}

		err = internal.SetUserEmail(selectedProfile.Email, false)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Credentials of profile %s set for current project.\n", selectedProfile.ProfileName)
	}
}

// CredentialsAlreadySet checks if the current repository already has the same credentials as the given profile.
// Returns true if both name and email match, false otherwise.
func CredentialsAlreadySet(profile models.ProfileConfig) bool {
	currentName, _ := internal.GetUserName()
	currentEmail, _ := internal.GetUserEmail()

	return profile.Name == currentName && profile.Email == currentEmail
}

func init() {
	rootCmd.AddCommand(initCmd)
}
