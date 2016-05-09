package config

import (
	"reflect"
	"testing"

	"github.com/bmizerany/assert"
)

func TestString(t *testing.T) {
	v, err := stringParser.Parse(reflect.TypeOf(""), "foo")
	assert.Equal(t, nil, err)
	assert.Equal(t, "foo", v)
}

func TestBool(t *testing.T) {
	cases := []struct {
		in  string
		out bool
		err bool
	}{
		{"false", false, false},
		{"true", true, false},
		{"yes", false, true},
	}

	for _, c := range cases {
		v, err := boolParser.Parse(reflect.TypeOf(false), c.in)
		assert.Equal(t, c.err, err != nil)
		assert.Equal(t, c.out, v)
	}
}

func TestInt(t *testing.T) {
	cases := []struct {
		in  string
		out interface{}
		err bool
	}{
		{"0", int(0), false},
		{"8", int8(8), false},
		{"16", int16(16), false},
		{"32", int32(32), false},
		{"64", int32(64), false},
		{"sa", int(0), true},
	}

	for _, c := range cases {
		v, err := intParser.Parse(reflect.TypeOf(c.out), c.in)
		assert.Equal(t, c.err, err != nil)
		assert.Equal(t, c.out, v)
	}
}
