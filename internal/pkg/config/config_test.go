package config_test

import (
	"testing"

	"github.com/haapjari/caliber/internal/pkg/config"
	"github.com/haapjari/caliber/internal/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	t.Parallel()
	log.NewLogger()

	tag, err := config.GetVersion()
	assert.NoError(t, err)

	t.Log(tag)
}
