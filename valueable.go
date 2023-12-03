package pqr

import (
	"reflect"
)

type Valueable interface {
	Values() []any
}

func newValueable[T any](t *T, cc []column) Valueable {
	if v, ok := any(t).(Valueable); ok {
		return v
	}
	return &valueable[T]{t, cc}
}

type valueable[T any] struct {
	v  *T
	cc []column
}

func (v *valueable[T]) Values() []any {
	el := reflect.ValueOf(v.v).Elem()
	pt := make([]any, len(v.cc))
	for _, c := range v.cc {
		// todo: add more validations here
		pt[c.index] = el.Field(c.index).Interface()
	}
	return pt
}
