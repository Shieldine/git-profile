package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/Shieldine/git-profile/internal"
)

// setupTestRepo creates a temporary directory and initializes a git repository in it.
func setupTestRepo(t *testing.T) (string, func()) {
	tempDir, err := os.MkdirTemp("", "gitrepo")
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		os.RemoveAll(tempDir)
		t.Fatal(err)
	}

	return tempDir, func() {
		os.RemoveAll(tempDir)
	}
}

func TestCheckGitRepo(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	// Change the working directory to the temporary git repository
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldDir)
	os.Chdir(tempDir)

	if !internal.CheckGitRepo() {
		t.Error("expected CheckGitRepo to return true in a git repository")
	}
}

func TestGetRepoOrigin(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	// Change the working directory to the temporary git repository
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldDir)
	os.Chdir(tempDir)

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

func TestSetUserName(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	// Change the working directory to the temporary git repository
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldDir)
	os.Chdir(tempDir)

	name := "Test User"
	if err := internal.SetUserName(name); err != nil {
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

func TestUnsetUserName(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	// Change the working directory to the temporary git repository
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldDir)
	os.Chdir(tempDir)

	cmd := exec.Command("git", "config", "--local", "user.name", "Test User")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	if err := internal.UnsetUserName(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cmd = exec.Command("git", "config", "--get", "--local", "user.name")
	if err := cmd.Run(); err == nil {
		t.Error("expected an error when getting unset user name, but got none")
	}
}

func TestGetUserName(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	// Change the working directory to the temporary git repository
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldDir)
	os.Chdir(tempDir)

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

func TestSetUserEmail(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	// Change the working directory to the temporary git repository
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldDir)
	os.Chdir(tempDir)

	email := "test@example.com"
	if err := internal.SetUserEmail(email); err != nil {
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

func TestUnsetUserEmail(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	// Change the working directory to the temporary git repository
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldDir)
	os.Chdir(tempDir)

	cmd := exec.Command("git", "config", "--local", "user.email", "test@example.com")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	if err := internal.UnsetUserEmail(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cmd = exec.Command("git", "config", "--get", "--local", "user.email")
	if err := cmd.Run(); err == nil {
		t.Error("expected an error when getting unset user email, but got none")
	}
}

func TestGetUserEmail(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	// Change the working directory to the temporary git repository
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldDir)
	os.Chdir(tempDir)

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
