package config

import (
	"os"
	"testing"

	"github.com/bmizerany/assert"
)

func TestLoad(t *testing.T) {
	var config struct {
		Name  string
		Debug bool
	}

	defer setEnv("Name", "prateek")()
	defer setEnv("Debug", "true")()

	err := Load(&config, nil, []Provider{EnvProvider})

	assert.Equal(t, nil, err)
	assert.Equal(t, "prateek", config.Name)
	assert.Equal(t, true, config.Debug)
}

func TestLoadDefaults(t *testing.T) {
	var config struct {
		Name  string
		Debug bool
	}

	err := Load(&config, nil, []Provider{EnvProvider})

	assert.Equal(t, nil, err)
	assert.Equal(t, "", config.Name)
	assert.Equal(t, false, config.Debug)
}

func setEnv(k, v string) func() {
	os.Setenv(k, v)
	return func() {
		os.Setenv(k, "")
	}
}
