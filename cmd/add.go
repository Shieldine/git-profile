/*
Copyright © 2024 Shieldine <EMAIL ADDRESS>
*/
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
	Use:     "add <profile-name>",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"a"},
	Short:   "Adds a new profile",
	Long: `Define a new profile with the short name <profile-name>. 

You will be asked to provide your credentials and an origin.
The origin of your current repository will already be filled in
and subject to confirm or change.`,
	Run: addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addRun(cmd *cobra.Command, args []string) {
	profileName := args[0]
	if (models.ProfileConfig{}) != internal.GetProfileByName(profileName) {
		fmt.Printf("Profile %s already exists\n", profileName)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("E-mail: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Signing key (enter to skip): ")
	signingKey, _ := reader.ReadString('\n')
	signingKey = strings.TrimSpace(signingKey)

	origin, _ := internal.GetRepoOrigin()
	fmt.Printf("Origin (enter to accept %s): ", origin)

	newOrigin, _ := reader.ReadString('\n')
	newOrigin = strings.TrimSpace(origin)

	if newOrigin == "" {
		newOrigin = origin
	}

	newProfile := models.ProfileConfig{
		ProfileName: profileName,
		Name:        name,
		Email:       email,
		SigningKey:  signingKey,
		Origin:      newOrigin,
	}

	err := internal.AddProfile(newProfile)
	if err != nil {
		fmt.Println("Error adding profile:", err)
		os.Exit(1)
	}

	fmt.Printf("Added profile: %s\n", profileName)
}
