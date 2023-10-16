// Package config contains the configuration for the application.
package config

import (
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/haapjari/caliber/internal/pkg/utils"
	"gopkg.in/yaml.v2"
)

// Config is the configuration for the application.
type Config struct {
	sync.Mutex

	configFilePath string

	Version string
	File    File
}

// File is the configuration file for the application.
type File struct {
	exclude []string
}

// New returns a new Config.
func New() (*Config, error) {
	c := &Config{
		Mutex: sync.Mutex{},
	}

	version, err := GetVersion()
	if err != nil {
		return nil, utils.HandleErr(err, "unable to retrieve version")
	}

	c.Version = version
	c.File = File{}

	configFilePath := GetConfigFilePath()

	if err = c.ReadConfigFile(configFilePath); err != nil {
		return nil, utils.HandleErr(err, "unable to read config file")
	}

	return c, err
}

// ReadConfigFile reads the configuration file.
//
//nolint:gosec
func (c *Config) ReadConfigFile(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return utils.HandleErr(err, "unable to read config file")
	}

	if err = yaml.Unmarshal(file, c.File); err != nil {
		return utils.HandleErr(err, "unable to unmarshal config file")
	}

	return nil
}

// GetVersion returns the version of the application.
// Function retrieves the tag from git and returns it.
//
//nolint:gosec
func GetVersion() (string, error) {
	revListCmd := exec.Command("git", "rev-list", "--tags", "--max-count=1")
	commitHash, err := revListCmd.Output()
	if err != nil {
		return "", utils.HandleErr(err, "failed to get latest commit hash")
	}

	commitHashStr := strings.TrimSpace(string(commitHash))
	describeCmd := exec.Command("git", "describe", "--tags", commitHashStr)
	tagName, err := describeCmd.Output()
	if err != nil {
		return "", utils.HandleErr(err, "failed to get the tag name")
	}

	trimmedTag := strings.TrimSpace(string(tagName))

	return trimmedTag, nil
}

// GetConfigFilePath returns the path to the configuration file.
// TODO
func GetConfigFilePath() string {
	return ".caliber.yml"
}

// GetDefaultFileName returns the default file name for the configuration file.
func GetDefaultFileName() string {
	return ".caliber.yml"
}
