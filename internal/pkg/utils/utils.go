// Package utils provides utilities for the application.
package utils

import (
	"fmt"

	"github.com/haapjari/caliber/internal/pkg/log"
)

// HandleErr handles errors and formats the error message with the given format
// string and arguments.
func HandleErr(err error, format string, args ...interface{}) error {
	errMsg := fmt.Sprintf(format, args...)
	log.Errorf("Error: %s: %v", errMsg, err)

	return fmt.Errorf("%s: %v", errMsg, err)
}
