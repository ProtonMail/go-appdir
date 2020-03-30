// +build !darwin,!windows

package appdir

import (
	"os"
	"path/filepath"
)

type dirs struct {
	name string
}

func (d *dirs) UserConfig() string {
	baseDir := os.Getenv("XDG_CONFIG_HOME")
	if baseDir == "" {
		baseDir = filepath.Join(os.Getenv("HOME"), ".config")
	}

	return filepath.Join(baseDir, d.name)
}

func (d *dirs) UserCache() string {
	baseDir := os.Getenv("XDG_CACHE_HOME")
	if baseDir == "" {
		baseDir = filepath.Join(os.Getenv("HOME"), ".cache")
	}

	return filepath.Join(baseDir, d.name)
}

func (d *dirs) UserData() string {
	baseDir := os.Getenv("XDG_DATA_HOME")
	if baseDir == "" {
		baseDir = filepath.Join(os.Getenv("HOME"), ".local", "share")
	}

	return filepath.Join(baseDir, d.name)
}

func (d *dirs) UserLogs() string {
	return filepath.Join(d.UserCache(), "logs")
}
