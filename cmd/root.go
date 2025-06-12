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
	"os"

	"github.com/spf13/cobra"
)

// version represents the current version of the git-profile tool
var version = "1.3.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "git-profile",
	Version: version,
	Short:   "Manage and automatically set git user profiles",
	Long: `git-profile is a simple CLI to manage and automatically set git user profiles based on the project's origin.

Save a profile together with its origin and let git-profile set the attributes next time you clone a new repository.
To make managing names and emails more convenient in general, git-profile offers further commands that will let you
check, unset and set credentials without creating a profile. You also get the option to do these things globally.

Common commands:
  git-profile init      - Automatically set credentials based on repository origin
  git-profile add       - Add a new profile
  git-profile list      - List all profiles
  git-profile update    - Update one or multiple existing profiles
  git-profile rm        - Remove a profile
  git-profile set       - Manually set a profile
  git-profile unset     - Unset git config
  git-profile check     - Check current credentials
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
