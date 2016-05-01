package config

import (
	"errors"
	"fmt"
	"reflect"
)

// Provider is the interface that wraps the Provide method.
type Provider interface {
	// Provide returns a string associated with the given key.
	Provide(k string) string
}

var ErrTypeNotSupported = errors.New("parser does not support given type")

// Parser is the interface that wraps the Parse method.
type Parser interface {
	// Parse parses the given value into the given type. Parse returns an error
	// ErrTypeNotSupported if it cannot parse values for the given type.
	Parse(t reflect.Type, v string) (interface{}, error)
}

// Validator is the interface that wraps the Validate method.
type Validator interface {
	// Validate validates and returns an error.
	Validate() error
}

func Load(ptr interface{}, userProvidedParsers []Parser, providers []Provider) error {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return errors.New("must provide a pointer")
	}

	v := reflect.ValueOf(ptr).Elem()
	if v.Kind() != reflect.Struct {
		return errors.New("must provide a pointer to struct")
	}

	parsers := append(builtInParsers, userProvidedParsers...)

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := fieldValue.Type()
		structField := t.Field(i)

		var providedValue string
		for _, provider := range providers {
			providedValue = provider.Provide(structField.Name)
			if providedValue != "" {
				break
			}
		}
		if providedValue == "" {
			continue
		}

		var parsedVal interface{}
		for _, parser := range parsers {
			val, err := parser.Parse(fieldType, providedValue)
			if err == ErrTypeNotSupported {
				continue
			}
			if err != nil {
				return fmt.Errorf("error converting %v to type %v: %v", providedValue, fieldValue.Kind(), err)
			}
			parsedVal = val
			break
		}
		fieldValue.Set(reflect.ValueOf(parsedVal))
	}

	if validator, ok := ptr.(Validator); ok {
		return validator.Validate()
	}

	return nil
}
