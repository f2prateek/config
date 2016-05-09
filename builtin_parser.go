package config

import (
	"reflect"
	"strconv"
)

var builtInParsers = []Parser{
	boolParser, uintParser, intParser, floatParser, stringParser,
}

type parserFunc func(t reflect.Type, v string) (interface{}, error)

func (f parserFunc) Parse(t reflect.Type, v string) (interface{}, error) {
	return f(t, v)
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

var intParser = parserFunc(func(t reflect.Type, v string) (interface{}, error) {
	switch t.Kind() {
	case reflect.Int:
		v, err := strconv.ParseInt(v, 0, 0)
		return int(v), err
	case reflect.Int8:
		v, err := strconv.ParseInt(v, 0, 8)
		return int8(v), err
	case reflect.Int16:
		v, err := strconv.ParseInt(v, 0, 16)
		return int16(v), err
	case reflect.Int32:
		v, err := strconv.ParseInt(v, 0, 32)
		return int32(v), err
	case reflect.Int64:
		return strconv.ParseInt(v, 0, 64)
	default:
		return nil, nil
	}
})

var uintParser = parserFunc(func(t reflect.Type, v string) (interface{}, error) {
	switch t.Kind() {
	case reflect.Uint:
		v, err := strconv.ParseUint(v, 0, 0)
		return uint(v), err
	case reflect.Uint8:
		v, err := strconv.ParseUint(v, 0, 8)
		return uint8(v), err
	case reflect.Uint16:
		v, err := strconv.ParseUint(v, 0, 16)
		return uint16(v), err
	case reflect.Uint32:
		v, err := strconv.ParseUint(v, 0, 32)
		return uint32(v), err
	case reflect.Uint64:
		return strconv.ParseUint(v, 0, 64)
	default:
		return nil, nil
	}
})

var floatParser = parserFunc(func(t reflect.Type, v string) (interface{}, error) {
	switch t.Kind() {
	case reflect.Float32:
		v, err := strconv.ParseFloat(v, 32)
		return float32(v), err
	case reflect.Float64:
		return strconv.ParseFloat(v, 64)
	default:
		return nil, nil
	}
})
