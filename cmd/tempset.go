/*
Copyright © 2024 Shieldine <74987363+Shieldine@users.noreply.github.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/Shieldine/git-profile/internal"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var tempSetCmd = &cobra.Command{
	Use:   "tempset",
	Short: "Set credentials without defining a profile",
	Long: `Set git credentials for the current repository without saving them in a profile.
The credentials can be passed as flags right away.
If you don't pass them, you will be asked to provide a name, email and signing key (optional).
`,
	Run: runTempSet,
}

func runTempSet(*cobra.Command, []string) {

	reader := bufio.NewReader(os.Stdin)

	if name == "" {
		currentName, err := internal.GetUserName()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Name (enter to keep %s): ", currentName)
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name != "" {
			err = internal.SetUserName(name)
			if err != nil {
				fmt.Printf("Error setting user name: %s\n", err)
				os.Exit(1)
			}
		}
	} else {
		err := internal.SetUserName(name)
		if err != nil {
			fmt.Printf("Error setting user name: %s\n", err)
			os.Exit(1)
		}
	}

	if email == "" {
		currentEmail, err := internal.GetUserEmail()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("E-mail (enter to keep %s): ", currentEmail)
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)

		if email != "" {
			err = internal.SetUserEmail(email)
			if err != nil {
				fmt.Printf("Error setting user email: %s\n", err)
				os.Exit(1)
			}
		}
	} else {
		err := internal.SetUserEmail(email)
		if err != nil {
			fmt.Printf("Error setting user email: %s\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("Credentials set successfully")
}

func init() {
	rootCmd.AddCommand(tempSetCmd)
	tempSetCmd.Flags().StringVarP(&name, "name", "n", "", "Pass the name directly")
	tempSetCmd.Flags().StringVarP(&email, "email", "e", "", "Pass the email directly")
}
