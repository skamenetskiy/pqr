package pqr

import (
	"strconv"
)

func newPlaceholders() *placeholders {
	return &placeholders{c: 1}
}

type placeholders struct {
	c int
	a []any
}

func (p *placeholders) next(args ...any) []string {
	r := make([]string, 0, len(args))
	for _, a := range args {
		p.a = append(p.a, a)
		r = append(r, "$"+strconv.Itoa(p.c))
		p.c++
	}
	return r
}

func (p *placeholders) nextAny(args ...any) []any {
	r := make([]any, 0, len(args))
	for _, a := range args {
		p.a = append(p.a, a)
		r = append(r, "$"+strconv.Itoa(p.c))
		p.c++
	}
	return r
}

func (p *placeholders) args() []any {
	return p.a
}
