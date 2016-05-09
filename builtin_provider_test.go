package config

import (
	"os"
	"testing"

	"github.com/bmizerany/assert"
)

func TestEnvProvider(t *testing.T) {
	defer setEnv("name", "prateek")()

	assert.Equal(t, "prateek", EnvProvider.Provide("name"))
	assert.Equal(t, "", EnvProvider.Provide("foo"))
}

func TestOSArgsProvider(t *testing.T) {
	defer setArgs("test", "--foo", "bar", "--qaz")()

	assert.Equal(t, "bar", OSArgsProvider.Provide("foo"))
	assert.Equal(t, "", OSArgsProvider.Provide("bar"))
	assert.Equal(t, "", OSArgsProvider.Provide("qaz"))
}

func setEnv(k, v string) func() {
	original := os.Getenv(k)
	os.Setenv(k, v)
	return func() {
		os.Setenv(k, original)
	}
}

func setArgs(args ...string) func() {
	original := os.Args
	os.Args = args
	return func() {
		os.Args = original
	}
}
