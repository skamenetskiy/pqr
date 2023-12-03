package pqr

import (
	"fmt"
	"slices"
)

// first element of slice, empty T otherwise.
func first[T any](pl []T) T {
	var p T
	if len(pl) > 0 {
		p = pl[0]
	}
	return p
}

func withoutIndex(in []any, index int) []any {
	r := make([]any, 0, len(in))
	for i, v := range in {
		if i == index {
			continue
		}
		r = append(r, v)
	}
	return r
}

func isZeroValue[K Key](k K) bool {
	switch v := any(k).(type) {
	case int32:
		return v == 0
	case uint32:
		return v == 0
	case int64:
		return v == 0
	case uint64:
		return v == 0
	case string:
		return v == ""
	}
	return false
}

type kv[K comparable, T any] struct {
	key   K
	value T
}

func mapToSortedSlice(m map[string]any) []kv[string, any] {
	pairs := make([]kv[string, any], 0, len(m))
	for k, v := range m {
		pairs = append(pairs, kv[string, any]{k, v})
	}
	slices.SortFunc[[]kv[string, any], kv[string, any]](pairs, func(a, b kv[string, any]) int {
		if a.key > b.key {
			return 1
		}
		return -1
	})
	return pairs
}

func isValidConditionType(t any) error {
	switch t.(type) {
	case int, int8, int16, int32, int64:
		return nil
	case uint, uint8, uint16, uint32, uint64:
		return nil
	case float32, float64:
		return nil
	case string:
		return nil
	case bool:
		return nil
	}
	return fmt.Errorf("condition %v is of invalid type %T", t, t)
}

func sliceToAny[T any](in []T) []any {
	r := make([]any, len(in))
	for i, v := range in {
		r[i] = v
	}
	return r
}
