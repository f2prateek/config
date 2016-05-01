package config

import "os"

var EnvProvider = providerFunc(func(s string) string {
	return os.Getenv(s)
})

type providerFunc func(string) string

func (f providerFunc) Provide(s string) string {
	return f(s)
}
