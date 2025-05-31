// Package test
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
package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/Shieldine/git-profile/internal"
)

// setupTestRepo creates a temporary directory and initializes a git repository in it.
// Returns the temporary directory path and a cleanup function to remove the directory.
func setupTestRepo(t *testing.T) (string, func()) {
	tempDir, err := os.MkdirTemp("", "gitrepo")
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		err := os.RemoveAll(tempDir)
		if err != nil {
			t.Fatal(err)
		}
		t.Fatal(err)
	}

	return tempDir, func() {
		err := os.RemoveAll(tempDir)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// TestCheckGitRepo tests the CheckGitRepo function to ensure it correctly identifies a git repository.
func TestCheckGitRepo(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func(dir string) {
		err := os.Chdir(dir)
		if err != nil {
			t.Fatal(err)
		}
	}(originalDir)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	if !internal.CheckGitRepo() {
		t.Error("expected CheckGitRepo to return true in a git repository")
	}
}

// TestGetRepoOrigin tests the GetRepoOrigin function to ensure it correctly extracts the hostname from a repository's origin URL.
func TestGetRepoOrigin(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func(dir string) {
		err := os.Chdir(dir)
		if err != nil {
			t.Fatal(err)
		}
	}(originalDir)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("git", "remote", "add", "origin", "https://example.com/repo.git")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	origin, err := internal.GetRepoOrigin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if origin != "example.com" {
		t.Errorf("expected origin to be 'example.com', got %s", origin)
	}
}

// TestSetUserNameLocal tests the SetUserName function with local scope to ensure it correctly sets the local user name.
func TestSetUserNameLocal(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func(dir string) {
		err := os.Chdir(dir)
		if err != nil {
			t.Fatal(err)
		}
	}(originalDir)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	name := "Test User"
	if err := internal.SetUserName(name, false); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cmd := exec.Command("git", "config", "--get", "--local", "user.name")
	output, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(string(output)) != name {
		t.Errorf("expected user name to be %s, got %s", name, output)
	}
}

// TestSetUserNameGlobal tests the SetUserName function with global scope to ensure it correctly sets the global username.
func TestSetUserNameGlobal(t *testing.T) {
	name := "Global Test User"
	if err := internal.SetUserName(name, true); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cmd := exec.Command("git", "config", "--get", "--global", "user.name")
	output, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(string(output)) != name {
		t.Errorf("expected global user name to be %s, got %s", name, output)
	}

	err = exec.Command("git", "config", "--global", "--unset", "user.name").Run()
	if err != nil {
		t.Errorf("unexpected error while unsetting global user name: %s", err)
	}
}

// TestUnsetUserNameLocal tests the UnsetUserName function with local scope to ensure it correctly removes the local username.
func TestUnsetUserNameLocal(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func(dir string) {
		err := os.Chdir(dir)
		if err != nil {
			t.Fatal(err)
		}
	}(originalDir)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("git", "config", "--local", "user.name", "Test User")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	if err := internal.UnsetUserName(false); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cmd = exec.Command("git", "config", "--get", "--local", "user.name")
	if err := cmd.Run(); err == nil {
		t.Error("expected an error when getting unset user name, but got none")
	}
}

// TestUnsetUserNameGlobal tests the UnsetUserName function with global scope to ensure it correctly removes the global user name.
func TestUnsetUserNameGlobal(t *testing.T) {
	cmd := exec.Command("git", "config", "--global", "user.name", "Global Test User")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	if err := internal.UnsetUserName(true); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cmd = exec.Command("git", "config", "--get", "--global", "user.name")
	if err := cmd.Run(); err == nil {
		t.Error("expected an error when getting unset global user name, but got none")
	}
}

// TestGetUserName tests the GetUserName function to ensure it correctly retrieves the local username.
func TestGetUserName(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func(dir string) {
		err := os.Chdir(dir)
		if err != nil {
			t.Fatal(err)
		}
	}(originalDir)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("git", "config", "--local", "user.name", "Test User")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	retrievedName, err := internal.GetUserName()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if retrievedName != "Test User" {
		t.Errorf("expected user name to be 'Test User', got %s", retrievedName)
	}
}

// TestGetGlobalUserName tests the GetGlobalUserName function to ensure it correctly retrieves the global user name.
func TestGetGlobalUserName(t *testing.T) {
	name := "Global Test User"
	cmd := exec.Command("git", "config", "--global", "user.name", name)
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	retrievedName, err := internal.GetGlobalUserName()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if retrievedName != name {
		t.Errorf("expected global user name to be '%s', got %s", name, retrievedName)
	}

	err = exec.Command("git", "config", "--global", "--unset", "user.name").Run()
	if err != nil {
		t.Errorf("unexpected error while unsetting global user name: %s", err)
	}
}

// TestSetUserEmailLocal tests the SetUserEmail function with local scope to ensure it correctly sets the local user email.
func TestSetUserEmailLocal(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func(dir string) {
		err := os.Chdir(dir)
		if err != nil {
			t.Fatal(err)
		}
	}(originalDir)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	email := "test@example.com"
	if err := internal.SetUserEmail(email, false); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cmd := exec.Command("git", "config", "--get", "--local", "user.email")
	output, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(string(output)) != email {
		t.Errorf("expected user email to be %s, got %s", email, output)
	}
}

// TestSetUserEmailGlobal tests the SetUserEmail function with global scope to ensure it correctly sets the global user email.
func TestSetUserEmailGlobal(t *testing.T) {
	email := "global@example.com"
	if err := internal.SetUserEmail(email, true); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cmd := exec.Command("git", "config", "--get", "--global", "user.email")
	output, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(string(output)) != email {
		t.Errorf("expected global user email to be %s, got %s", email, output)
	}

	err = exec.Command("git", "config", "--global", "--unset", "user.email").Run()
	if err != nil {
		t.Errorf("unexpected error while unsetting global user email: %s", err)
	}
}

// TestUnsetUserEmailLocal tests the UnsetUserEmail function with local scope to ensure it correctly removes the local user email.
func TestUnsetUserEmailLocal(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func(dir string) {
		err := os.Chdir(dir)
		if err != nil {
			t.Fatal(err)
		}
	}(originalDir)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("git", "config", "--local", "user.email", "test@example.com")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	if err := internal.UnsetUserEmail(false); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cmd = exec.Command("git", "config", "--get", "--local", "user.email")
	if err := cmd.Run(); err == nil {
		t.Error("expected an error when getting unset user email, but got none")
	}
}

// TestUnsetUserEmailGlobal tests the UnsetUserEmail function with global scope to ensure it correctly removes the global user email.
func TestUnsetUserEmailGlobal(t *testing.T) {
	cmd := exec.Command("git", "config", "--global", "user.email", "global@example.com")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	if err := internal.UnsetUserEmail(true); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cmd = exec.Command("git", "config", "--get", "--global", "user.email")
	if err := cmd.Run(); err == nil {
		t.Error("expected an error when getting unset global user email, but got none")
	}
}

// TestGetUserEmail tests the GetUserEmail function to ensure it correctly retrieves the local user email.
func TestGetUserEmail(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func(dir string) {
		err := os.Chdir(dir)
		if err != nil {
			t.Fatal(err)
		}
	}(originalDir)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("git", "config", "--local", "user.email", "test@example.com")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	retrievedEmail, err := internal.GetUserEmail()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if retrievedEmail != "test@example.com" {
		t.Errorf("expected user email to be 'test@example.com', got %s", retrievedEmail)
	}
}

// TestGetGlobalUserEmail tests the GetGlobalUserEmail function to ensure it correctly retrieves the global user email.
func TestGetGlobalUserEmail(t *testing.T) {
	email := "global@example.com"
	cmd := exec.Command("git", "config", "--global", "user.email", email)
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	retrievedEmail, err := internal.GetGlobalUserEmail()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if retrievedEmail != email {
		t.Errorf("expected global user email to be '%s', got %s", email, retrievedEmail)
	}

	err = exec.Command("git", "config", "--global", "--unset", "user.email").Run()
	if err != nil {
		t.Errorf("unexpected error while unsetting global user email: %s", err)
	}
}
