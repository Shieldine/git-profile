// Package cmd provides command-line functionality for git-profile
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

var (
	newName   string
	newEmail  string
	newOrigin string
	oldName   string
	oldEmail  string
	oldOrigin string
)

// editCmd represents the update command
var editCmd = &cobra.Command{
	Use:     "update [profile-name]",
	Aliases: []string{"edit", "u", "e"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "Update one or multiple profiles",
	Long: `Update profiles based on provided criteria.

When a profile name is provided, updates only that specific profile.
Without a profile name, updates all profiles matching the filter criteria.

Examples:
  # Update a specific profile interactively
  git-profile update myprofile

  # Update a specific profile with flags
  git-profile update myprofile --name "New Name" --email "new@example.com"

  # Update all profiles with a specific email
  git-profile update --old-email old@example.com --email new@example.com

  # Update all profiles with a specific name
  git-profile update --old-name "Old Name" --name "New Name"

  # Update all profiles with a specific origin
  git-profile update --old-origin github.com --origin gitlab.com
`,
	Run: runUpdate,
}

// runUpdate handles the update command execution.
// It supports two modes of operation:
// 1. Single profile update: When a profile name is provided as an argument
// 2. Batch update: When no profile name is provided, but filter criteria are specified
//
// In single profile mode, the user can update a profile interactively or using flags.
// In batch mode, the command updates all profiles matching the filter criteria.
func runUpdate(_ *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)

	// Single profile update
	if len(args) == 1 {
		profileName := args[0]
		oldProfile := internal.GetProfileByName(profileName)

		if (models.ProfileConfig{}) == oldProfile {
			fmt.Printf("Profile %s doesn't exist.\n", profileName)
			return
		}

		if newName == "" {
			fmt.Printf("Name (enter to keep %s): ", oldProfile.Name)
			newName, _ = reader.ReadString('\n')
			newName = strings.TrimSpace(newName)

			if newName == "" {
				newName = oldProfile.Name
			}
		}

		if newEmail == "" {
			fmt.Printf("E-mail (enter to keep %s): ", oldProfile.Email)
			newEmail, _ = reader.ReadString('\n')
			newEmail = strings.TrimSpace(newEmail)

			if newEmail == "" {
				newEmail = oldProfile.Email
			}
		}

		currentOrigin, _ := internal.GetRepoOrigin()

		if newOrigin == "" {
			fmt.Printf("Origin (enter to keep %s): ", oldProfile.Origin)

			newOrigin, _ = reader.ReadString('\n')
			newOrigin = strings.TrimSpace(newOrigin)

			if newOrigin == "" {
				newOrigin = oldProfile.Origin
			}
		} else if newOrigin == "auto" {
			newOrigin = currentOrigin
		}

		err := internal.EditProfile(profileName, models.ProfileConfig{
			ProfileName: profileName,
			Name:        newName,
			Email:       newEmail,
			Origin:      newOrigin,
		})
		if err != nil {
			fmt.Printf("Error updating profile: %v\n", err)
			return
		}

		fmt.Printf("Profile %s updated\n", profileName)
		return
	}

	// Batch update
	if oldName == "" && oldEmail == "" && oldOrigin == "" {
		fmt.Println("Error: When updating multiple profiles, you must specify at least one filter criteria (--old-name, --old-email, or --old-origin).")
		return
	}

	if newName == "" && newEmail == "" && newOrigin == "" {
		fmt.Println("Error: When updating multiple profiles, you must specify at least one new value (--name, --email, or --origin).")
		return
	}

	// Get all profiles
	profiles := internal.GetAllProfiles()
	if len(profiles) == 0 {
		fmt.Println("No profiles to update.")
		return
	}

	// Filter and update profiles
	updatedCount := 0
	for _, profile := range profiles {

		if (oldName != "" && profile.Name != oldName) ||
			(oldEmail != "" && profile.Email != oldEmail) ||
			(oldOrigin != "" && profile.Origin != oldOrigin) {
			continue
		}

		updatedProfile := models.ProfileConfig{
			ProfileName: profile.ProfileName,
			Name:        profile.Name,
			Email:       profile.Email,
			Origin:      profile.Origin,
		}

		if newName != "" {
			updatedProfile.Name = newName
		}
		if newEmail != "" {
			updatedProfile.Email = newEmail
		}
		if newOrigin != "" {
			if newOrigin == "auto" {
				currentOrigin, _ := internal.GetRepoOrigin()
				updatedProfile.Origin = currentOrigin
			} else {
				updatedProfile.Origin = newOrigin
			}
		}

		err := internal.EditProfile(profile.ProfileName, updatedProfile)
		if err != nil {
			fmt.Printf("Error updating profile %s: %v\n", profile.ProfileName, err)
			continue
		}

		fmt.Printf("Profile %s updated\n", profile.ProfileName)
		updatedCount++
	}

	if updatedCount > 0 {
		fmt.Printf("\nSuccessfully updated %d profile(s).\n", updatedCount)
	} else {
		fmt.Println("No profiles matched the filter criteria.")
	}
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringVarP(&newName, "name", "n", "", "Set the new name value")
	editCmd.Flags().StringVarP(&newEmail, "email", "e", "", "Set the new email value")
	editCmd.Flags().StringVarP(&newOrigin, "origin", "o", "", "Set the new origin value. Type \"auto\" to use current repository's origin")

	editCmd.Flags().StringVar(&oldName, "old-name", "", "Filter profiles by name")
	editCmd.Flags().StringVar(&oldEmail, "old-email", "", "Filter profiles by email")
	editCmd.Flags().StringVar(&oldOrigin, "old-origin", "", "Filter profiles by origin")
}
