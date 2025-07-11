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

// unsetCmd represents the unset command for removing Git attributes
var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Reset attribute config to none",
	Long: `Resets git attributes for current repository or globally.
If you unset local config, git will default to your global config.
If you unset global config, git will have no default credentials.

Examples:
  # Unset local repository attributes
  git-profile unset

  # Unset global attributes
  git-profile unset --global
`,
	Run: runUnset,
}

// runUnset executes the unset command logic.
// It removes Git user configuration (name and email) from either local repository or global scope.
// If --global flag is used, it unsets the global Git configuration; otherwise, it unsets the local repository configuration.
func runUnset(cmd *cobra.Command, _ []string) {
	global, _ := cmd.Flags().GetBool("global")

	if global {
		fmt.Println("warning: removing global git credentials")
	} else {
		if !internal.CheckGitRepo() {
			fmt.Println("error: not a git repository")
			os.Exit(1)
		}

		fmt.Println("warning: git will default to global credentials without local configuration")
	}

	err := internal.UnsetUserName(global)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	err = internal.UnsetUserEmail(global)
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}
}

func init() {
	unsetCmd.Flags().BoolP("global", "g", false, "Unset the credentials globally instead of for the current repository")

	rootCmd.AddCommand(unsetCmd)
}
