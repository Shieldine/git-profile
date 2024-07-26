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

func CheckGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

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

func SetUserName(name string) error {
	if !CheckGitRepo() {
		return errors.New("not a git repository")
	}
	cmd := exec.Command("git", "config", "user.name", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func UnsetUserName() error {
	if !CheckGitRepo() {
		return errors.New("not a git repository")
	}

	_, err := GetUserName()

	if err != nil {
		return errors.New("no local username to unset")
	}

	cmd := exec.Command("git", "config", "--unset", "user.name")
	_, err = cmd.Output()

	if err != nil {
		return err
	}
	return nil
}

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

func SetUserEmail(email string) error {
	if !CheckGitRepo() {
		return errors.New("not a git repository")
	}
	cmd := exec.Command("git", "config", "user.email", email)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

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

func UnsetUserEmail() error {
	if !CheckGitRepo() {
		return errors.New("not a git repository")
	}

	_, err := GetUserEmail()

	if err != nil {
		return errors.New("no local email to unset")
	}

	cmd := exec.Command("git", "config", "--unset", "user.email")
	_, err = cmd.Output()

	if err != nil {
		return err
	}
	return nil
}
