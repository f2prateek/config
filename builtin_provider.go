package config

import "os"

var EnvProvider = providerFunc(func(s string) string {
	return os.Getenv(s)
})

var ArgsProvider = providerFunc(func(s string) string {
	for i, arg := range os.Args {
		if arg != "--"+s {
			continue
		}
		if i == len(os.Args)-1 {
			return ""
		}
		return os.Args[i+1]
	}
	return ""
})

type providerFunc func(string) string

func (f providerFunc) Provide(s string) string {
	return f(s)
}
