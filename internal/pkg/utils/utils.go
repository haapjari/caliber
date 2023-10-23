// Package utils provides utilities for the application.
package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/haapjari/caliber/internal/pkg/log"
)

// HandleErr handles errors and formats the error message with the given format
// string and arguments.
func HandleErr(err error, format string, args ...interface{}) error {
	errMsg := fmt.Sprintf(format, args...)
	log.Errorf("Error: %s: %v", errMsg, err)

	return fmt.Errorf("%s: %v", errMsg, err)
}

// ReturnAbsolutePath returns absolute path of the passed "path" variable,
// if that is a valid path, otherwise return error.
func ReturnAbsolutePath(path string) (string, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		log.Errorf("unable to read absolute path: %w", err)
		return "", err
	}

	_, err = os.Stat(absolutePath)
	if err != nil {
		log.Errorf("invalid path: %v", err)
		return "", err
	}

	return absolutePath, nil
}
