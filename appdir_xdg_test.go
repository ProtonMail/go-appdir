package appdir

import (
	"os"
	"testing"
)

func TestUserConfig(t *testing.T) {
	if err := os.Setenv("HOME", "/home/user"); err != nil {
		panic(err)
	}

	for _, tc := range [...]struct {
		name          string
		overrideValue string
		expected      string
	}{
		{
			name:     "default config directory",
			expected: "/home/user/.config/app",
		},
		{
			name:          "config directory set to absolute",
			overrideValue: "/config",
			expected:      "/config/app",
		},
		{
			name:          "config directory set to relative",
			overrideValue: "config",
			expected:      "config/app",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d := dirs{name: "app"}

			if err := os.Setenv("XDG_CONFIG_HOME", tc.overrideValue); err != nil {
				panic(err)
			}

			if have := d.UserConfig(); have != tc.expected {
				t.Errorf("expected %q, found %q", tc.expected, have)
			}
		})
	}
}

func TestUserCache(t *testing.T) {
	if err := os.Setenv("HOME", "/home/user"); err != nil {
		panic(err)
	}

	for _, tc := range [...]struct {
		name          string
		overrideValue string
		expected      string
	}{
		{
			name:     "default cache directory",
			expected: "/home/user/.cache/app",
		},
		{
			name:          "cache directory set to absolute",
			overrideValue: "/cache",
			expected:      "/cache/app",
		},
		{
			name:          "cache directory set to relative",
			overrideValue: "cache",
			expected:      "cache/app",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d := dirs{name: "app"}

			if err := os.Setenv("XDG_CACHE_HOME", tc.overrideValue); err != nil {
				panic(err)
			}

			if have := d.UserCache(); have != tc.expected {
				t.Errorf("expected %q, found %q", tc.expected, have)
			}
		})
	}
}

func TestUserLogs(t *testing.T) {
	if err := os.Setenv("HOME", "/home/user"); err != nil {
		panic(err)
	}

	for _, tc := range [...]struct {
		name          string
		overrideValue string
		expected      string
	}{
		{
			name:     "default logs directory",
			expected: "/home/user/.local/state/app",
		},
		{
			name:          "logs directory set to absolute",
			overrideValue: "/logs",
			expected:      "/logs/app",
		},
		{
			name:          "logs directory set to relative",
			overrideValue: "logs",
			expected:      "logs/app",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d := dirs{name: "app"}

			if err := os.Setenv("XDG_STATE_HOME", tc.overrideValue); err != nil {
				panic(err)
			}

			if have := d.UserLogs(); have != tc.expected {
				t.Errorf("expected %q, found %q", tc.expected, have)
			}
		})
	}
}

func TestUserData(t *testing.T) {
	if err := os.Setenv("HOME", "/home/user"); err != nil {
		panic(err)
	}

	for _, tc := range [...]struct {
		name          string
		overrideValue string
		expected      string
	}{
		{
			name:     "default data directory",
			expected: "/home/user/.local/share/app",
		},
		{
			name:          "data directory set to absolute",
			overrideValue: "/data",
			expected:      "/data/app",
		},
		{
			name:          "data directory set to relative",
			overrideValue: "data",
			expected:      "data/app",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d := dirs{name: "app"}

			if err := os.Setenv("XDG_DATA_HOME", tc.overrideValue); err != nil {
				panic(err)
			}

			if have := d.UserData(); have != tc.expected {
				t.Errorf("expected %q, found %q", tc.expected, have)
			}
		})
	}
}
