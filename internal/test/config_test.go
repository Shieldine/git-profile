package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/Shieldine/git-profile/internal"
	"github.com/Shieldine/git-profile/models"
)

func setupTempConfig(t *testing.T) (string, func()) {
	tempDir, err := os.MkdirTemp("", "configTest")
	if err != nil {
		t.Fatal(err)
	}

	tempConfigPath := filepath.Join(tempDir, "config.toml")
	internal.Conf = internal.Config{Profiles: []models.ProfileConfig{}}
	internal.SetConfigPath(tempConfigPath)

	err = internal.SaveConfig()
	if err != nil {
		t.Fatal(err)
	}

	return tempConfigPath, func() {
		err := os.RemoveAll(tempDir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
}

func TestLoadConfig(t *testing.T) {
	_, cleanup := setupTempConfig(t)
	defer cleanup()

	err := internal.LoadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(internal.Conf.Profiles) != 0 {
		t.Errorf("expected empty profiles, got %d", len(internal.Conf.Profiles))
	}
}

func TestSaveConfig(t *testing.T) {
	configPath, cleanup := setupTempConfig(t)
	defer cleanup()

	profile := models.ProfileConfig{ProfileName: "test"}
	internal.Conf.Profiles = append(internal.Conf.Profiles, profile)

	err := internal.SaveConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var loadedConfig internal.Config
	_, err = toml.DecodeFile(configPath, &loadedConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(loadedConfig.Profiles) != 1 || loadedConfig.Profiles[0].ProfileName != "test" {
		t.Errorf("expected profile name 'test', got %v", loadedConfig.Profiles)
	}
}

func TestAddProfile(t *testing.T) {
	_, cleanup := setupTempConfig(t)
	defer cleanup()

	profile := models.ProfileConfig{ProfileName: "test"}
	err := internal.AddProfile(profile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(internal.Conf.Profiles) != 1 || internal.Conf.Profiles[0].ProfileName != "test" {
		t.Errorf("expected profile name 'test', got %v", internal.Conf.Profiles)
	}
}

func TestEditProfile(t *testing.T) {
	_, cleanup := setupTempConfig(t)
	defer cleanup()

	profile := models.ProfileConfig{ProfileName: "test"}
	err := internal.AddProfile(profile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updatedProfile := models.ProfileConfig{ProfileName: "updated"}
	err = internal.EditProfile("test", updatedProfile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(internal.Conf.Profiles) != 1 || internal.Conf.Profiles[0].ProfileName != "updated" {
		t.Errorf("expected profile name 'updated', got %v", internal.Conf.Profiles)
	}
}

func TestDeleteProfile(t *testing.T) {
	_, cleanup := setupTempConfig(t)
	defer cleanup()

	profile := models.ProfileConfig{ProfileName: "test"}
	err := internal.AddProfile(profile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = internal.DeleteProfile("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(internal.Conf.Profiles) != 0 {
		t.Errorf("expected empty profiles, got %d", len(internal.Conf.Profiles))
	}
}

func TestGetProfileByName(t *testing.T) {
	_, cleanup := setupTempConfig(t)
	defer cleanup()

	profile := models.ProfileConfig{ProfileName: "test"}
	err := internal.AddProfile(profile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	retrievedProfile := internal.GetProfileByName("test")
	if retrievedProfile.ProfileName != "test" {
		t.Errorf("expected profile name 'test', got %s", retrievedProfile.ProfileName)
	}
}

func TestGetAllProfiles(t *testing.T) {
	_, cleanup := setupTempConfig(t)
	defer cleanup()

	profile1 := models.ProfileConfig{ProfileName: "test1"}
	profile2 := models.ProfileConfig{ProfileName: "test2"}
	err := internal.AddProfile(profile1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = internal.AddProfile(profile2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	profiles := internal.GetAllProfiles()
	if len(profiles) != 2 {
		t.Errorf("expected 2 profiles, got %d", len(profiles))
	}
}

func TestClearConfig(t *testing.T) {
	_, cleanup := setupTempConfig(t)
	defer cleanup()

	profile := models.ProfileConfig{ProfileName: "test"}
	err := internal.AddProfile(profile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = internal.ClearConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = internal.LoadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	profiles := internal.Conf.Profiles

	if len(profiles) != 0 {
		t.Errorf("expected empty profiles, got %d", len(internal.Conf.Profiles))
	}
}

func TestGetProfilesByOrigin(t *testing.T) {
	_, cleanup := setupTempConfig(t)
	defer cleanup()

	profile1 := models.ProfileConfig{ProfileName: "test1", Origin: "origin1"}
	profile2 := models.ProfileConfig{ProfileName: "test2", Origin: "origin2"}
	profile3 := models.ProfileConfig{ProfileName: "test3", Origin: "origin1"}
	err := internal.AddProfile(profile1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = internal.AddProfile(profile2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = internal.AddProfile(profile3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	profiles := internal.GetProfilesByOrigin("origin1")
	if len(profiles) != 2 {
		t.Errorf("expected 2 profiles, got %d", len(profiles))
	}
}
