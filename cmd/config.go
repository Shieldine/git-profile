/*
Copyright © 2024 Shieldine <74987363+Shieldine@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Shieldine/git-profile/internal"
	"github.com/spf13/cobra"
)

var (
	editorChoice string
)

var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"c"},
	Short:   "Edit profile configuration file",
	Long: `Open and edit the config file containing all profiles.
You can manually type in new profiles by using the following scheme:

[[profiles]]
  profile_name = ""
  name = ""
  email = ""
  origin = ""
`,
	Run: runConfig,
}

func runConfig(*cobra.Command, []string) {
	editor := editorChoice
	if editor == "" {
		editor = "vim"
	}

	editorCmd := exec.Command(editor, internal.GetConfigPath())
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	if err := editorCmd.Run(); err != nil {
		fmt.Printf("Failed to open editor: %v\n", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVarP(&editorChoice, "editor", "e", "", "Specify the editor to use (e.g. nano, code). Vim will be used as default")
}
