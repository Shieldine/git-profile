// Package internal
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
package internal

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/Shieldine/git-profile/custom_errors"
)

// CheckGitRepo checks if the current directory is inside a Git repository.
// Returns true if inside a Git repository, false otherwise.
func CheckGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// GetRepoOrigin retrieves the origin URL of the Git repository and extracts the hostname.
// Returns the hostname (e.g., "github.com") from the remote origin URL.
// Returns an error if not in a Git repository or if the origin URL cannot be retrieved.
func GetRepoOrigin() (string, error) {
	if !CheckGitRepo() {
		return "", errors.New("not a git repository")
	}
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	originURL := strings.TrimSpace(string(output))

	if strings.HasPrefix(originURL, "https://") {
		originURL = strings.TrimPrefix(originURL, "https://")
	} else if strings.HasPrefix(originURL, "http://") {
		originURL = strings.TrimPrefix(originURL, "http://")
	} else if strings.HasPrefix(originURL, "git@") {
		originURL = strings.TrimPrefix(originURL, "git@")
		if idx := strings.Index(originURL, ":"); idx != -1 {
			originURL = originURL[:idx]
		}
	}

	if idx := strings.Index(originURL, "/"); idx != -1 {
		originURL = originURL[:idx]
	}

	return originURL, nil
}

// SetUserName sets the Git user.name configuration.
// If global is true, sets the global configuration; otherwise sets local repository configuration.
// Returns an error if not in a Git repository (when global is false) or if the git command fails.
func SetUserName(name string, global bool) error {
	if !global && !CheckGitRepo() {
		return errors.New("not a git repository")
	}

	args := []string{"config", "user.name", name}
	if global {
		args = []string{"config", "--global", "user.name", name}
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// UnsetUserName removes the Git user.name configuration.
// If global is true, unsets the global configuration; otherwise unsets local repository configuration.
// Returns an error if not in a Git repository (when global is false), if no username is set, or if the git command fails.
func UnsetUserName(global bool) error {
	if !global && !CheckGitRepo() {
		return errors.New("not a git repository")
	}

	_, err := GetUserName()
	if err != nil {
		return errors.New("no local username to unset")
	}

	args := []string{"config", "--unset", "user.name"}
	if global {
		args = []string{"config", "--global", "--unset", "user.name"}
	}

	cmd := exec.Command("git", args...)
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

// GetUserName retrieves the local Git user.name configuration.
// Returns the username string or an error if not in a Git repository or if the username is not set.
// Returns a custom NotSetError if the username is not configured locally.
func GetUserName() (string, error) {
	if !CheckGitRepo() {
		return "", errors.New("not a git repository")
	}
	cmd := exec.Command("git", "config", "--get", "--local", "user.name")
	output, err := cmd.CombinedOutput()

	if err != nil {
		var exitError *exec.ExitError

		ok := errors.As(err, &exitError)
		if ok && exitError.ExitCode() == 1 && err.Error() == "exit status 1" {
			return "", &custom_errors.NotSetError{ConfigName: "username"}
		} else {
			return "", err
		}
	}
	return strings.TrimSpace(string(output)), nil
}

// SetUserEmail sets the Git user.email configuration.
// If global is true, sets the global configuration; otherwise sets local repository configuration.
// Returns an error if not in a Git repository (when global is false) or if the git command fails.
func SetUserEmail(email string, global bool) error {
	if !global && !CheckGitRepo() {
		return errors.New("not a git repository")
	}

	args := []string{"config", "user.email", email}
	if global {
		args = []string{"config", "--global", "user.email", email}
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// GetUserEmail retrieves the local Git user.email configuration.
// Returns the email string or an error if not in a Git repository or if the email is not set.
// Returns a custom NotSetError if the email is not configured locally.
func GetUserEmail() (string, error) {
	if !CheckGitRepo() {
		return "", errors.New("not a git repository")
	}
	cmd := exec.Command("git", "config", "--get", "--local", "user.email")
	output, err := cmd.CombinedOutput()

	if err != nil {
		var exitError *exec.ExitError

		ok := errors.As(err, &exitError)
		if ok && exitError.ExitCode() == 1 && err.Error() == "exit status 1" {
			return "", &custom_errors.NotSetError{ConfigName: "email"}
		} else {
			return "", err
		}
	}
	return strings.TrimSpace(string(output)), nil
}

// GetGlobalUserName retrieves the global Git user.name configuration.
// Returns the username string or an error if the username is not set globally.
// Returns a custom NotSetError if the username is not configured globally.
func GetGlobalUserName() (string, error) {
	cmd := exec.Command("git", "config", "--get", "--global", "user.name")
	output, err := cmd.CombinedOutput()

	if err != nil {
		var exitError *exec.ExitError

		ok := errors.As(err, &exitError)
		if ok && exitError.ExitCode() == 1 && err.Error() == "exit status 1" {
			return "", &custom_errors.NotSetError{ConfigName: "global username"}
		} else {
			return "", err
		}
	}
	return strings.TrimSpace(string(output)), nil
}

// GetGlobalUserEmail retrieves the global Git user.email configuration.
// Returns the email string or an error if the email is not set globally.
// Returns a custom NotSetError if the email is not configured globally.
func GetGlobalUserEmail() (string, error) {
	cmd := exec.Command("git", "config", "--get", "--global", "user.email")
	output, err := cmd.CombinedOutput()

	if err != nil {
		var exitError *exec.ExitError

		ok := errors.As(err, &exitError)
		if ok && exitError.ExitCode() == 1 && err.Error() == "exit status 1" {
			return "", &custom_errors.NotSetError{ConfigName: "global email"}
		} else {
			return "", err
		}
	}
	return strings.TrimSpace(string(output)), nil
}

// UnsetUserEmail removes the Git user.email configuration.
// If global is true, unsets the global configuration; otherwise unsets local repository configuration.
// Returns an error if not in a Git repository (when global is false), if no email is set, or if the git command fails.
func UnsetUserEmail(global bool) error {
	if !global && !CheckGitRepo() {
		return errors.New("not a git repository")
	}

	_, err := GetUserEmail()
	if err != nil {
		return errors.New("no local email to unset")
	}

	args := []string{"config", "--unset", "user.email"}
	if global {
		args = []string{"config", "--global", "--unset", "user.email"}
	}

	cmd := exec.Command("git", args...)
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	return nil
}
