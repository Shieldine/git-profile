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

// setCmd represents the set command for changing Git profiles
var setCmd = &cobra.Command{
	Use:     "set <profile-name>",
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(1),
	Short:   "Set profile for current repository or globally",
	Long:    `Change the current repository's profile to <profile-name>, or set it globally with --global flag.`,
	Run:     runSet,
}

// runSet executes the set command logic.
// It sets the Git user configuration (name and email) based on the specified profile.
// If --global flag is used, it sets the global Git configuration; otherwise, it sets the local repository configuration.
func runSet(cmd *cobra.Command, args []string) {
	global, _ := cmd.Flags().GetBool("global")

	if !global && !internal.CheckGitRepo() {
		fmt.Println("error: not a git repository")
		os.Exit(1)
	}

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

	if !global {
		currentOrigin, err := internal.GetRepoOrigin()
		if err != nil {
			fmt.Printf("Error getting repository origin: %s\n", err)
			os.Exit(1)
		}

		if profile.Origin != currentOrigin {
			fmt.Println("warning: profile origin and repo origin don't match.")
			fmt.Printf("	Repo origin: %s\n", currentOrigin)
			fmt.Printf("	Profile origin: %s\n", profile.Origin)
			fmt.Println()
		}
	}

	var currentName, currentEmail string
	var nameErr, emailErr error

	if global {
		currentName, nameErr = internal.GetGlobalUserName()
		currentEmail, emailErr = internal.GetGlobalUserEmail()
	} else {
		currentName, nameErr = internal.GetUserName()
		currentEmail, emailErr = internal.GetUserEmail()
	}

	if nameErr != nil {
		var notSetErr *custom_errors.NotSetError
		if !errors.As(nameErr, &notSetErr) {
			fmt.Println("error: ", nameErr)
			os.Exit(1)
		}
	}

	if emailErr != nil {
		var notSetErr *custom_errors.NotSetError
		if !errors.As(emailErr, &notSetErr) {
			fmt.Println("error: ", emailErr)
			os.Exit(1)
		}
	}

	if profile.Name == currentName && profile.Email == currentEmail {
		if global {
			fmt.Println("Global configuration already has correct credentials. Nothing to do.")
		} else {
			fmt.Println("Repository already has correct credentials. Nothing to do.")
		}
		return
	}

	err := internal.SetUserName(profile.Name, global)
	if err != nil {
		fmt.Printf("Error setting user name: %s\n", err)
		os.Exit(1)
	}

	err = internal.SetUserEmail(profile.Email, global)
	if err != nil {
		fmt.Printf("Error setting user email: %s\n", err)
		os.Exit(1)
	}

	if global {
		fmt.Printf("Profile %s set globally.\n", profileName)
	} else {
		fmt.Printf("Profile %s set for current repository.\n", profileName)
	}
}

// ReadAnswer prompts the user for a yes/no answer and validates the input.
// It continues to prompt until a valid answer ('y' or 'n') is provided.
// Returns the validated answer as a lowercase string.
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

// init initializes the set command and adds it to the root command.
// It also defines the --global/-g flag for setting Git configuration globally.
func init() {
	setCmd.Flags().BoolP("global", "g", false, "Set the profile globally instead of for the current repository")

	rootCmd.AddCommand(setCmd)
}
