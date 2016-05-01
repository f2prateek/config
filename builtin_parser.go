package config

import (
	"reflect"
	"strconv"
)

// kindParser only supports parsing types of kind k.
type kindParser struct {
	k reflect.Kind
	f func(v string) (interface{}, error)
}

func (p *kindParser) Parse(t reflect.Type, v string) (interface{}, error) {
	if t.Kind() != p.k {
		return nil, ErrTypeNotSupported
	}

	return p.f(v)
}

var boolParser = &kindParser{
	k: reflect.Bool,
	f: func(v string) (interface{}, error) {
		return strconv.ParseBool(v)
	},
}

var stringParser = &kindParser{
	k: reflect.String,
	f: func(v string) (interface{}, error) {
		return v, nil
	},
}
