package models

type ProfileConfig struct {
	ProfileName string `toml:"profile_name"`
	Name        string `toml:"name"`
	Email       string `toml:"email"`
	SigningKey  string `toml:"signing_key"`
	Origin      string `toml:"origin"`
}
