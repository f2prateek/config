package config

import (
	"reflect"
	"strconv"
)

var builtInParsers = []Parser{
	boolParser, stringParser,
}

// kindParser only supports parsing types of kind k.
type kindParser struct {
	k reflect.Kind
	f func(v string) (interface{}, error)
}

func (p *kindParser) Parse(t reflect.Type, v string) (interface{}, error) {
	if t.Kind() != p.k {
		return nil, nil
	}

	return p.f(v)
}

func newKindParser(k reflect.Kind, f func(v string) (interface{}, error)) Parser {
	return &kindParser{
		k: k,
		f: f,
	}
}

var boolParser = newKindParser(reflect.Bool, func(v string) (interface{}, error) {
	return strconv.ParseBool(v)
})

var stringParser = newKindParser(reflect.String, func(v string) (interface{}, error) {
	return v, nil
})
