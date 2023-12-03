package pqr

import (
	"reflect"
)

type Pointable interface {
	Pointers() []any
}

func newPointable[T any](t *T, cc []column) Pointable {
	if p, ok := any(t).(Pointable); ok {
		return p
	}
	return &pointable[T]{t, cc}
}

type pointable[T any] struct {
	v  *T
	cc []column
}

func (p *pointable[T]) Pointers() []any {
	v := reflect.ValueOf(p.v).Elem()
	pt := make([]any, len(p.cc))
	for _, c := range p.cc {
		// todo: add more validations here
		pt[c.index] = v.Field(c.index).Addr().Interface()
	}
	return pt
}
