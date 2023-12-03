package pqr

import (
	"reflect"
	"testing"
)

func Test_first(t *testing.T) {
	type args[T any] struct {
		pl []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{"one", args[int]{[]int{1}}, 1},
		{"two", args[int]{[]int{1, 2, 3}}, 1},
		{"three", args[int]{[]int{3, 2, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := first(tt.args.pl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("first() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidConditionType(t *testing.T) {
	type args struct {
		t any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"int", args{int(1)}, false},
		{"int8", args{int8(1)}, false},
		{"int16", args{int16(1)}, false},
		{"int32", args{int32(1)}, false},
		{"int64", args{int64(1)}, false},
		{"uint", args{uint(1)}, false},
		{"uint8", args{uint8(1)}, false},
		{"uint16", args{uint16(1)}, false},
		{"uint32", args{uint32(1)}, false},
		{"uint64", args{uint64(1)}, false},
		{"float32", args{float32(1)}, false},
		{"float64", args{float64(1)}, false},
		{"string", args{"1"}, false},
		{"bool", args{true}, false},
		{"bool", args{false}, false},
		{"struct", args{struct{}{}}, true},
		{"func", args{func() {}}, true},
		{"slice", args{[]int{1, 2, 3}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isValidConditionType(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("isValidConditionType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sliceToAny(t *testing.T) {
	type args[T any] struct {
		in []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []any
	}
	tests := []testCase[int]{
		{"one", args[int]{[]int{1}}, []any{1}},
		{"many", args[int]{[]int{1, 2, 3}}, []any{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceToAny(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

type isZeroValueTestCase[K Key] struct {
	name string
	arg  K
	want bool
}

func testIsZeroValue[K Key](t *testing.T, tests []isZeroValueTestCase[K]) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isZeroValue(tt.arg); got != tt.want {
				t.Errorf("isZeroValue(%T) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func Test_isZeroValueInt32(t *testing.T) {
	testIsZeroValue[int32](t, []isZeroValueTestCase[int32]{
		{"not zero", int32(1), false},
		{"is zero", int32(0), true},
	})
}

func Test_isZeroValueInt64(t *testing.T) {
	testIsZeroValue[int64](t, []isZeroValueTestCase[int64]{
		{"not zero", int64(1), false},
		{"is zero", int64(0), true},
	})
}

func Test_isZeroValueUint32(t *testing.T) {
	testIsZeroValue[uint32](t, []isZeroValueTestCase[uint32]{
		{"not zero", uint32(1), false},
		{"is zero", uint32(0), true},
	})
}

func Test_isZeroValueUint64(t *testing.T) {
	testIsZeroValue[uint64](t, []isZeroValueTestCase[uint64]{
		{"not zero", uint64(1), false},
		{"is zero", uint64(0), true},
	})
}

func Test_isZeroValueString(t *testing.T) {
	testIsZeroValue[string](t, []isZeroValueTestCase[string]{
		{"not zero", "not zero", false},
		{"is zero", "", true},
	})
}

func Test_mapToSortedSlice(t *testing.T) {
	type args struct {
		m map[string]any
	}
	tests := []struct {
		name string
		args args
		want []kv[string, any]
	}{
		{"sorted", args{map[string]any{"a": 1, "b": 2, "c": 3}}, []kv[string, any]{{"a", 1}, {"b", 2}, {"c", 3}}},
		{"unsorted", args{map[string]any{"c": 3, "b": 2, "a": 1}}, []kv[string, any]{{"a", 1}, {"b", 2}, {"c", 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapToSortedSlice(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapToSortedSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_withoutIndex(t *testing.T) {
	type args struct {
		in    []any
		index int
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{"0", args{[]any{1, 2, 3, 4, 5}, 0}, []any{2, 3, 4, 5}},
		{"1", args{[]any{1, 2, 3, 4, 5}, 1}, []any{1, 3, 4, 5}},
		{"2", args{[]any{1, 2, 3, 4, 5}, 2}, []any{1, 2, 4, 5}},
		{"3", args{[]any{1, 2, 3, 4, 5}, 3}, []any{1, 2, 3, 5}},
		{"4", args{[]any{1, 2, 3, 4, 5}, 4}, []any{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := withoutIndex(tt.args.in, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("withoutIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
