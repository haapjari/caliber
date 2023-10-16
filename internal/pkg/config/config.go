// Package config contains the configuration for the application.
package config

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/haapjari/caliber/internal/pkg/utils"
)

// Config is the configuration for the application.
type Config struct {
	sync.Mutex

	version string
}

// New returns a new Config.
func New() (*Config, error) {
	version, err := GetVersion()
	if err != nil {
		return nil, utils.HandleErr(err, "unable to retrieve version")
	}

	return &Config{
		version: version,
		Mutex:   sync.Mutex{},
	}, err
}

// GetVersion returns the version of the application.
// Function retrieves the tag from git and returns it.
func GetVersion() (string, error) {
	revListCmd := exec.Command("git", "rev-list", "--tags", "--max-count=1")
	commitHash, err := revListCmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get latest commit hash: %w", err)
	}

	// 2. Describe that commit to get the tag name
	describeCmd := exec.Command("git", "describe", "--tags", strings.TrimSpace(string(commitHash)))
	tagName, err := describeCmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get tag name: %w", err)
	}

	return strings.TrimSpace(string(tagName)), nil
}
