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
	if os.Getenv("XDG_CONFIG_HOME") != "" {
		return os.Getenv("XDG_CONFIG_HOME")
	}

	return filepath.Join(os.Getenv("HOME"), ".config", d.name)
}

func (d *dirs) UserCache() string {
	if os.Getenv("XDG_CACHE_HOME") != "" {
		return os.Getenv("XDG_CACHE_HOME")
	}

	return filepath.Join(os.Getenv("HOME"), ".cache", d.name)
}
