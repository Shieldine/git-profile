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
	"github.com/spf13/cobra"
	"os"
)

// checkCmd represents the check command for displaying current git credentials
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Display the currently set attributes",
	Long: `Check what attributes are currently set in the current project or globally.

This command displays the name and email currently configured in git.
Use the --global flag to check the global git configuration instead of the local repository configuration.

Examples:
  # Check local repository attributes
  git-profile check

  # Check global attributes
  git-profile check --global
`,
	Run: runCheck,
}

// runCheck executes the check command logic.
// It displays the current Git user configuration (name and email).
// If --global flag is used, it shows the global Git configuration; otherwise, it shows the local repository configuration.
func runCheck(cmd *cobra.Command, _ []string) {
	global, _ := cmd.Flags().GetBool("global")

	var name, email string
	var nameErr, emailErr error

	if global {
		name, nameErr = internal.GetGlobalUserName()
		email, emailErr = internal.GetGlobalUserEmail()
		fmt.Println("Global Git Configuration:")
	} else {
		if !internal.CheckGitRepo() {
			fmt.Println("error: not a git repository")
			os.Exit(1)
		}

		name, nameErr = internal.GetUserName()
		email, emailErr = internal.GetUserEmail()
		fmt.Println("Local Git Configuration:")
	}

	if nameErr != nil {
		fmt.Printf("error: %v\n", nameErr)
	} else {
		fmt.Printf("Current name: %s\n", name)
	}

	if emailErr != nil {
		fmt.Printf("error: %v\n", emailErr)
	} else {
		fmt.Printf("Current email: %s\n", email)
	}
}

func init() {
	checkCmd.Flags().BoolP("global", "g", false, "Check the global credentials instead of the current repository")

	rootCmd.AddCommand(checkCmd)
}
