# config

Package config provides a pluggable mechanism for loading configurations. It accepts a struct and takes of retrieving arguments, parsing them, and validating the final result.

```go
func Load(ptr interface{}, parsers []Parser, providers []Provider) error
```
