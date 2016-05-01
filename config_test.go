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

	os.Setenv("Name", "prateek")
	os.Setenv("Debug", "true")

	err := Load(&config, []Parser{boolParser, stringParser}, []Provider{EnvProvider})

	assert.Equal(t, nil, err)
	assert.Equal(t, "prateek", config.Name)
	assert.Equal(t, true, config.Debug)
}
