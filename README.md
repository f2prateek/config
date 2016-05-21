# config

Package config provides a pluggable mechanism for loading configurations. It accepts a struct and takes of retrieving arguments, parsing them, and validating the final result.

# Usage

```go
func Load(ptr interface{}, parsers []Parser, providers []Provider) error
```

# Example

```go
type AppConfig struct {
  Name  string
  Version int
  Debug bool
}

func (c *AppConfig) Validate() error() {
  if c.Name == "" {
    return errors.New("Name is required")
  }
  return nil
}

err := config.Load(&appConfig, nil, []Provider{config.EnvProvider, config.ArgsProvider})

fmt.Println(appConfig, err)
```

`Debug=false go run main.go --Name foo --Version bar`
