package config

import (
	"errors"
	"reflect"
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

func TestParserError(t *testing.T) {
	var config struct {
		Debug bool
	}

	badParser := parserFunc(func(t reflect.Type, v string) (interface{}, error) {
		return nil, errors.New("test")
	})

	err := Load(&config, []Parser{badParser}, testProvider(map[string]string{
		"Debug": "foo",
	}))

	assert.Equal(t, errors.New("error converting \"foo\" to type bool: test"), err)
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
