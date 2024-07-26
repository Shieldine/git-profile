package models

type ProfileConfig struct {
	ProfileName string `toml:"profile_name"`
	Name        string `toml:"name"`
	Email       string `toml:"email"`
	Origin      string `toml:"origin"`
}
