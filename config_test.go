package config

import (
	"errors"
	"testing"

	"github.com/bmizerany/assert"
)

func testProvider(m map[string]string) []Provider {
	return []Provider{providerFunc(func(s string) string {
		return m[s]
	})}
}

func TestLoad(t *testing.T) {
	var config struct {
		Name  string
		Debug bool
	}

	err := Load(&config, nil, testProvider(map[string]string{
		"Name":  "prateek",
		"Debug": "true",
	}))

	assert.Equal(t, nil, err)
	assert.Equal(t, "prateek", config.Name)
	assert.Equal(t, true, config.Debug)
}

func TestLoadDefaults(t *testing.T) {
	var config struct {
		Name  string
		Debug bool
	}

	err := Load(&config, nil, testProvider(map[string]string{}))

	assert.Equal(t, nil, err)
	assert.Equal(t, "", config.Name)
	assert.Equal(t, false, config.Debug)
}

type validatingConfig struct {
	Foo string
}

func (v *validatingConfig) Validate() error {
	return errors.New("test")
}

func TestLoadValidates(t *testing.T) {
	var config validatingConfig

	err := Load(&config, nil, testProvider(map[string]string{}))

	assert.Equal(t, errors.New("test"), err)
}
