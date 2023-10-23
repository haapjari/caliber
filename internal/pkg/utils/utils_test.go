package utils_test

import (
	"testing"

	"github.com/haapjari/caliber/internal/pkg/log"
	"github.com/haapjari/caliber/internal/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestReturnAbsolutePath(t *testing.T) {
	t.Parallel()
	log.NewLogger()

	_, err := utils.ReturnAbsolutePath(".")
	assert.NoError(t, err, "should not error, valid path")

	_, err = utils.ReturnAbsolutePath("invalid")
	assert.Error(t, err, "should error, invalid path")
}
