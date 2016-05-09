# config

Package config provides a pluggable mechanism for loading configurations. It accepts a struct and takes of retrieving arguments, parsing them, and validating the final result.

```go
func Load(ptr interface{}, parsers []Parser, providers []Provider) error
```

```go
type AppConfig struct {
  Name  string
  Version int
  Debug bool
}

err := config.Load(&appConfig, nil, []Provider{config.EnvProvider, config.ArgsProvider})

fmt.Println(appConfig, err)
```
