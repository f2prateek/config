package config

import (
	"os"
	"testing"

	"github.com/bmizerany/assert"
)

func TestEnvProvider(t *testing.T) {
	os.Setenv("name", "prateek")
	defer os.Setenv("name", "")

	assert.Equal(t, "prateek", EnvProvider.Provide("name"))
	assert.Equal(t, "", EnvProvider.Provide("foo"))
}
