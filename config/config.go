package config

import (
	"os"
	"runtime"
)

// Config basic global config
type Config struct {
	HomeDir string
}

var c *Config

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func init() {
	c = new(Config)
	c.HomeDir = userHomeDir()
}

// Get config
func Get() *Config {
	return c
}
