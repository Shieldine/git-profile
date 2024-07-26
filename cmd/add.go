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
	"os"
	"strings"

	"github.com/Shieldine/git-profile/internal"
	"github.com/Shieldine/git-profile/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [profile-name]",
	Args:    cobra.MaximumNArgs(1),
	Aliases: []string{"a"},
	Short:   "Add a new profile",
	Long: `Define a new profile with the short name <profile-name>.
Passing the profile name as an arg is optional. If not provided, you will
be asked to provide one.

You will be asked to provide your credentials and an origin.
Use flags to provide them directly.
The origin of your current repository will already be filled in
and subject to confirm or change.`,
	Run: runAdd,
}

func runAdd(_ *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)

	if len(args) == 0 {
		fmt.Print("Short name of the profile: ")
		profileName, _ = reader.ReadString('\n')
		profileName = strings.TrimSpace(profileName)
	} else {
		profileName = args[0]
	}

	if (models.ProfileConfig{}) != internal.GetProfileByName(profileName) {
		fmt.Printf("Profile %s already exists\n", profileName)
		return
	}

	if name == "" {
		fmt.Print("Name: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)
	}

	if email == "" {
		fmt.Print("E-mail: ")
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)
	}

	currentOrigin, _ := internal.GetRepoOrigin()
	newOrigin := ""

	if origin == "" {
		fmt.Printf("Origin (enter to accept %s): ", currentOrigin)

		newOrigin, _ = reader.ReadString('\n')
		newOrigin = strings.TrimSpace(newOrigin)

		if newOrigin == "" {
			newOrigin = currentOrigin
		}
	} else {
		if origin == "auto" {
			newOrigin = currentOrigin
		} else {
			newOrigin = origin
		}
	}

	newProfile := models.ProfileConfig{
		ProfileName: profileName,
		Name:        name,
		Email:       email,
		Origin:      newOrigin,
	}

	err := internal.AddProfile(newProfile)
	if err != nil {
		fmt.Println("Error adding profile:", err)
		os.Exit(1)
	}

	fmt.Printf("Added profile: %s for origin %s\n", profileName, newOrigin)
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Set the name directly")
	addCmd.Flags().StringVarP(&email, "email", "e", "", "Set the email directly")
	addCmd.Flags().StringVarP(&origin, "origin", "o", "", "Set the origin directly."+
		" Type \"auto\" to accept origin of the current repository")
}
