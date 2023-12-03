package pqr

import (
	"fmt"
	"strings"
)

// Condition interface describes the search condition in Find like methods.
type Condition interface {
	isNil() bool
	isSlice() bool
	toSQL(p *placeholders) (string, error)
}

type In[T any] []T

func (c In[T]) isNil() bool {
	return len(c) == 0
}

func (c In[T]) isSlice() bool {
	return true
}

func (c In[T]) toSQL(p *placeholders) (string, error) {
	s := p.next(sliceToAny[T](c)...)
	return strings.Join(s, ", "), nil
}

// And condition.
type And []Condition

func (c And) isNil() bool {
	return len(c) == 0
}

// isSlice is always false. Although it is actually a slice,
// the logic in toSQL method covers the required behaviour.
func (c And) isSlice() bool {
	return false
}

func (c And) toSQL(p *placeholders) (string, error) {
	return toAndOrSQL(c, " AND ", p)
}

type Or []Condition

func (c Or) isNil() bool {
	return len(c) == 0
}

func (c Or) isSlice() bool {
	return false
}

func (c Or) toSQL(p *placeholders) (string, error) {
	return toAndOrSQL(c, " OR ", p)
}

type Eq map[string]any

func (c Eq) isNil() bool {
	return len(c) == 0
}

func (c Eq) isSlice() bool {
	return false
}

func (c Eq) toSQL(p *placeholders) (string, error) {
	return toEqNqSQL(c, true, p)
}

type Nq map[string]any

func (c Nq) isNil() bool {
	return len(c) == 0
}

func (c Nq) isSlice() bool {
	return false
}

func (c Nq) toSQL(p *placeholders) (string, error) {
	return toEqNqSQL(c, false, p)
}

type Lt map[string]any

func (c Lt) isNil() bool {
	return len(c) == 0
}

func (c Lt) isSlice() bool {
	return false
}

func (c Lt) toSQL(p *placeholders) (string, error) {
	return toLtGtSQL(c, "<", p)
}

type Lte map[string]any

func (c Lte) isNil() bool {
	return len(c) == 0
}

func (c Lte) isSlice() bool {
	return false
}

func (c Lte) toSQL(p *placeholders) (string, error) {
	return toLtGtSQL(c, "<=", p)
}

type Gt map[string]any

func (c Gt) isNil() bool {
	return len(c) == 0
}

func (c Gt) isSlice() bool {
	return false
}

func (c Gt) toSQL(p *placeholders) (string, error) {
	return toLtGtSQL(c, ">", p)
}

type Gte map[string]any

func (c Gte) isNil() bool {
	return len(c) == 0
}

func (c Gte) isSlice() bool {
	return false
}

func (c Gte) toSQL(p *placeholders) (string, error) {
	return toLtGtSQL(c, ">=", p)
}

func toEqNqSQL(c map[string]any, o bool, p *placeholders) (string, error) {
	pairs := mapToSortedSlice(c)
	s := make([]string, 0, len(c))
	for _, v := range pairs {
		if sv, ok := v.value.(Condition); ok {
			if sv.isSlice() {
				inS, err := sv.toSQL(p)
				if err != nil {
					return "", err
				}
				op := "IN"
				if !o {
					op = "NOT IN"
				}
				s = append(s, fmt.Sprintf(`"%s" %s (%s)`, v.key, op, inS))
			} else {
				es, err := sv.toSQL(p)
				if err != nil {
					return "", err
				}
				s = append(s, es)
			}
			continue
		}
		op := "="
		if !o {
			op = "<>"
		}
		s = append(s, fmt.Sprintf(`"%s" %s %s`, append([]any{v.key, op}, p.nextAny(v.value)...)...))
	}
	return strings.Join(s, " AND "), nil
}

func toAndOrSQL(c []Condition, o string, p *placeholders) (string, error) {
	r := make([]string, len(c))
	for i, v := range c {
		s, err := v.toSQL(p)
		if err != nil {
			return "", err
		}
		r[i] = `(` + s + `)`
	}
	return strings.Join(r, o), nil
}

func toLtGtSQL(c map[string]any, o string, p *placeholders) (string, error) {
	var (
		s   = make([]string, 0, len(c))
		err error
	)
	pairs := mapToSortedSlice(c)
	for _, v := range pairs {
		if err = isValidConditionType(v.value); err != nil {
			return "", err
		}
		s = append(s, fmt.Sprintf(`"%s" %s %s`, append([]any{v.key, o}, p.nextAny(v.value)...)...))
	}
	return strings.Join(s, " AND "), nil
}
