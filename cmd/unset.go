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
)

var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Reset credential config to none",
	Long: `Resets git attributes for current repository.
If you do this, git will default to your global config.`,
	Run: runUnset,
}

func runUnset(*cobra.Command, []string) {

	fmt.Println("warning: git will default to global credentials without local configuration")

	err := internal.UnsetUserName()
	if err != nil {
		fmt.Println(err)
	}

	err = internal.UnsetUserEmail()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Process complete.")
}

func init() {
	rootCmd.AddCommand(unsetCmd)
}
