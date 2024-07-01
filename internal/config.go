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
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/Shieldine/git-profile/models"
)

type Config struct {
	Profiles []models.ProfileConfig `toml:"profiles"`
}

var (
	Conf       Config
	configPath string
)

func init() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error determining executable path:", err)
		os.Exit(1)
	}
	configPath = filepath.Join(filepath.Dir(exePath), "config.toml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		file, err := os.Create(configPath)
		if err != nil {
			fmt.Printf("Failed to create config file: %v\n", err)
			os.Exit(1)
		}
		_ = file.Close()
		Conf = Config{Profiles: []models.ProfileConfig{}}
	}

	err = LoadConfig()
	if err != nil {
		fmt.Println("Error loading config file:", err)
		os.Exit(1)
	}
}

func SetConfigPath(path string) {
	configPath = path
}

func LoadConfig() error {
	Conf.Profiles = []models.ProfileConfig{}
	if _, err := toml.DecodeFile(configPath, &Conf); err != nil {
		return fmt.Errorf("failed to decode internal file: %v", err)
	}
	return nil
}

func SaveConfig() error {
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("failed to save config file: %v", err)
	}

	if err := toml.NewEncoder(file).Encode(Conf); err != nil {
		return fmt.Errorf("failed to encode config to file: %v", err)
	}

	_ = file.Close()
	return nil
}

func AddProfile(profile models.ProfileConfig) error {
	for _, existingProfile := range Conf.Profiles {
		if existingProfile.ProfileName == profile.ProfileName {
			return fmt.Errorf("profile with name %s already exists", profile.ProfileName)
		}
	}

	Conf.Profiles = append(Conf.Profiles, profile)
	return SaveConfig()
}

func EditProfile(profileName string, updatedProfile models.ProfileConfig) error {
	for i, existingProfile := range Conf.Profiles {
		if existingProfile.ProfileName == profileName {
			Conf.Profiles[i] = updatedProfile
			return SaveConfig()
		}
	}
	return fmt.Errorf("profile with name %s not found", profileName)
}

func DeleteProfile(profileName string) error {
	for i, existingProfile := range Conf.Profiles {
		if existingProfile.ProfileName == profileName {
			Conf.Profiles = append(Conf.Profiles[:i], Conf.Profiles[i+1:]...)
			return SaveConfig()
		}
	}
	return fmt.Errorf("profile with name %s not found", profileName)
}

func GetProfileByName(profileName string) models.ProfileConfig {
	for _, existingProfile := range Conf.Profiles {
		if existingProfile.ProfileName == profileName {
			return existingProfile
		}
	}
	return models.ProfileConfig{}
}

func GetAllProfiles() []models.ProfileConfig {
	return Conf.Profiles
}

func GetConfigPath() string {
	return configPath
}

func ClearConfig() error {
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("failed to reset config file: %v", err)
	}
	_ = file.Close()
	return nil
}

func GetProfilesByOrigin(origin string) []models.ProfileConfig {
	var profiles []models.ProfileConfig

	for _, profile := range Conf.Profiles {
		if profile.Origin == origin {
			profiles = append(profiles, profile)
		}
	}

	return profiles
}
