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
	"errors"
	"fmt"
	"github.com/Shieldine/git-profile/custom_errors"
	"os"
	"strings"

	"github.com/Shieldine/git-profile/internal"
	"github.com/Shieldine/git-profile/models"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:     "set <profile-name>",
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(1),
	Short:   "Set profile for current repository",
	Long:    `Change the current repository's profile to <profile-name>.'`,
	Run:     runSet,
}

func runSet(cmd *cobra.Command, args []string) {

	profileName := args[0]
	profile := internal.GetProfileByName(profileName)

	if (models.ProfileConfig{}) == profile {
		fmt.Printf("Profile %s doesn't exist.\n", profileName)
		fmt.Print("Would you like to create it? (y/n): ")

		answer := ReadAnswer()

		if answer == "n" {
			fmt.Println("Nothing to do.")
			return
		}

		if answer == "y" {
			runAdd(cmd, []string{profileName})
		}
	}

	profile = internal.GetProfileByName(profileName)

	currentOrigin, _ := internal.GetRepoOrigin()

	if profile.Origin != currentOrigin {
		fmt.Println("warning: profile origin and repo origin don't match.")
		fmt.Printf("	Repo origin: %s\n", currentOrigin)
		fmt.Printf("	Profile origin: %s\n", profile.Origin)
		fmt.Println()
	}

	currentName, err := internal.GetUserName()
	currentEmail, _ := internal.GetUserEmail()

	if err != nil {
		var notSetErr *custom_errors.NotSetError
		if !errors.As(err, &notSetErr) {
			fmt.Println("error: ", err)
			os.Exit(1)
		}
	}

	if profile.Name == currentName && profile.Email == currentEmail {
		fmt.Println("Repository already has correct credentials. Nothing to do.")
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

	fmt.Printf("Profile %s set for current repository.\n", profileName)
}

func ReadAnswer() string {
	reader := bufio.NewReader(os.Stdin)
	answer := ""

	for {
		answer, _ = reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)

		if answer == "n" {
			break
		} else if answer == "y" {
			break
		} else {
			fmt.Println("Invalid choice. Choices are (y/n):")
		}
	}
	return answer
}

func init() {
	rootCmd.AddCommand(setCmd)
}
