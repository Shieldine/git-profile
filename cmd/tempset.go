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
	"os"
	"strings"

	"github.com/Shieldine/git-profile/custom_errors"
	"github.com/Shieldine/git-profile/internal"
	"github.com/spf13/cobra"
)

var tempSetCmd = &cobra.Command{
	Use:   "tempset",
	Short: "Set credentials without defining a profile",
	Long: `
Set git credentials for the current repository or globally without saving them in a profile.
The credentials can be passed as flags right away.
If you don't pass them, you will be asked to provide a name and an email.
`,
	Run: runTempSet,
}

// runTempSet executes the tempset command logic.
// It sets Git user configuration (name and email) without creating a profile.
// If --global flag is used, it sets the global Git configuration; otherwise, it sets the local repository configuration.
// Credentials can be provided via flags or will be prompted interactively.
func runTempSet(cmd *cobra.Command, _ []string) {
	reader := bufio.NewReader(os.Stdin)
	global, _ := cmd.Flags().GetBool("global")

	if name == "" {
		var currentName string
		var err error

		if global {
			currentName, err = internal.GetGlobalUserName()
		} else {
			currentName, err = internal.GetUserName()
		}

		if err != nil {
			var notSetErr *custom_errors.NotSetError

			if !errors.As(err, &notSetErr) {
				fmt.Printf("error: %v\n", err)
				os.Exit(1)
			}
		}

		if currentName != "" {
			fmt.Printf("Name (enter to keep %s): ", currentName)
		} else {
			fmt.Print("Name: ")
		}

		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name != "" {
			err = internal.SetUserName(name, global)
			if err != nil {
				fmt.Printf("Error while setting user name: %s\n", err)
				os.Exit(1)
			}
		}
	} else {
		err := internal.SetUserName(name, global)

		if err != nil {
			fmt.Printf("Error while setting user name: %s\n", err)
			os.Exit(1)
		}
	}

	if email == "" {
		var currentEmail string
		var err error

		if global {
			currentEmail, err = internal.GetGlobalUserEmail()
		} else {
			currentEmail, err = internal.GetUserEmail()
		}

		if err != nil {
			var notSetErr *custom_errors.NotSetError

			if !errors.As(err, &notSetErr) {
				fmt.Printf("error: %v\n", err)
				os.Exit(1)
			}
		}

		if currentEmail != "" {
			fmt.Printf("Email (enter to keep %s): ", currentEmail)
		} else {
			fmt.Print("E-Mail: ")
		}
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)

		if email != "" {
			err = internal.SetUserEmail(email, global)
			if err != nil {
				fmt.Printf("Error while setting user email: %s\n", err)
				os.Exit(1)
			}
		}
	} else {
		err := internal.SetUserEmail(email, global)
		if err != nil {
			fmt.Printf("Error while setting user email: %s\n", err)
			os.Exit(1)
		}
	}

	if global {
		fmt.Println("Global credentials set successfully")
	} else {
		fmt.Println("Credentials set successfully")
	}
}

// init initializes the tempset command and adds it to the root command.
// It defines flags for name, email, and global scope configuration.
func init() {
	rootCmd.AddCommand(tempSetCmd)
	tempSetCmd.Flags().StringVarP(&name, "name", "n", "", "Pass the name directly")
	tempSetCmd.Flags().StringVarP(&email, "email", "e", "", "Pass the email directly")
	tempSetCmd.Flags().BoolP("global", "g", false, "Set the credentials globally instead of for the current repository")
}
