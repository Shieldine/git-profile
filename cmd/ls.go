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
	"fmt"
	"github.com/Shieldine/git-profile/internal"
	"github.com/Shieldine/git-profile/models"
	"github.com/spf13/cobra"
)

var (
	profileName string
	name        string
	email       string
	origin      string
)

// lsCmd represents the list command for displaying git profiles
var lsCmd = &cobra.Command{
	Use:     "list [profile-name]",
	Aliases: []string{"l", "ls"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "List profiles",
	Long: `Display profiles currently present in your config.

Provide a profile name to list the attributes of the specified profile.
Use flags to filter for a specific origin, name or email.

Examples:
  # List all profiles
  git-profile list

  # Show details of a specific profile
  git-profile list myprofile

  # List all profiles with a specific email
  git-profile list --email user@example.com

  # List all profiles with a specific name
  git-profile list --name "John Doe"

  # List all profiles with a specific origin
  git-profile list --origin github.com
`,
	Run: runLs,
}

// runLs handles the list command execution.
// It supports two modes of operation:
// 1. Display a specific profile by name (when an argument is provided)
// 2. List all profiles, optionally filtered by name, email, or origin
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

// PrintProfile formats and prints the details of a Git profile.
// It displays the profile name, origin, name, and email in a readable format.
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
