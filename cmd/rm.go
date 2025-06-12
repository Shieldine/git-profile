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
	"os"
)

var all bool

// rmCmd represents the remove command for deleting git profiles
var rmCmd = &cobra.Command{
	Use:   "rm [profile-name]",
	Short: "Remove existing profiles",
	Args:  cobra.MaximumNArgs(1),
	Long: `Remove one or multiple profiles from the configuration.

Use --all flag to remove all profiles.
Use other flags to remove all profiles containing a specific name, email or origin.

Provide <profile-name> to remove only the profile called <profile-name>.
<profile-name> and filtering flags cannot be provided together.

This action cannot be undone.

Examples:
  # Remove a specific profile
  git-profile rm myprofile

  # Remove all profiles
  git-profile rm --all

  # Remove all profiles with a specific email
  git-profile rm --email user@example.com

  # Remove all profiles with a specific name
  git-profile rm --name "John Doe"

  # Remove all profiles with a specific origin
  git-profile rm --origin github.com
`,
	Run: runRm,
}

// runRm handles the remove command execution.
// It supports three modes of operation:
// 1. Remove all profiles (--all flag)
// 2. Remove a specific profile by name (argument)
// 3. Remove profiles matching filter criteria (--name, --email, --origin flags)
func runRm(_ *cobra.Command, args []string) {
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
}

func init() {
	rootCmd.AddCommand(rmCmd)
	rmCmd.Flags().BoolVarP(&all, "all", "a", false, "Remove all profiles")
	rmCmd.Flags().StringVarP(&name, "name", "n", "", "Remove profiles with name")
	rmCmd.Flags().StringVarP(&email, "email", "e", "", "Remove profiles with email")
	rmCmd.Flags().StringVarP(&origin, "origin", "o", "", "Remove profiles with origin")
}
