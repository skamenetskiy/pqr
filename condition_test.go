package pqr

import (
	"errors"
	"testing"
)

func TestAnd_isNil(t *testing.T) {
	tests := []struct {
		name string
		c    And
		want bool
	}{
		{"true", And{}, true},
		{"false", And{Eq{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isNil(); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnd_isSlice(t *testing.T) {
	tests := []struct {
		name string
		c    And
		want bool
	}{
		{" false", And{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isSlice(); got != tt.want {
				t.Errorf("isSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnd_toSQL(t *testing.T) {
	type args struct {
		p *placeholders
	}
	tests := []struct {
		name    string
		c       And
		args    args
		want    string
		wantErr bool
	}{
		{"simple", And{Eq{"a": 1}}, args{newPlaceholders()}, `("a" = $1)`, false},
		{"simple", And{Eq{"a": 1, "b": 2}}, args{newPlaceholders()}, `("a" = $1 AND "b" = $2)`, false},
		{"advanced", And{Eq{"a": 1, "b": 2}, Eq{"c": 3, "d": 4}}, args{newPlaceholders()}, `("a" = $1 AND "b" = $2) AND ("c" = $3 AND "d" = $4)`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.toSQL(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("toSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEq_isNil(t *testing.T) {
	tests := []struct {
		name string
		c    Eq
		want bool
	}{
		{"true", Eq{}, true},
		{"false", Eq{"a": 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isNil(); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEq_isSlice(t *testing.T) {
	tests := []struct {
		name string
		c    Eq
		want bool
	}{
		{"false", Eq{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isSlice(); got != tt.want {
				t.Errorf("isSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEq_toSQL(t *testing.T) {
	type args struct {
		p *placeholders
	}
	tests := []struct {
		name    string
		c       Eq
		args    args
		want    string
		wantErr bool
	}{
		{"simple", Eq{"a": 1}, args{newPlaceholders()}, `"a" = $1`, false},
		{"simple string", Eq{"a": "1"}, args{newPlaceholders()}, `"a" = $1`, false},
		{"advanced", Eq{"a": 1, "b": 2}, args{newPlaceholders()}, `"a" = $1 AND "b" = $2`, false},
		{"in", Eq{"a": In[int]{1, 2, 3}}, args{newPlaceholders()}, `"a" IN ($1, $2, $3)`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.toSQL(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("toSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGt_isNil(t *testing.T) {
	tests := []struct {
		name string
		c    Gt
		want bool
	}{
		{"true", Gt{}, true},
		{"false", Gt{"a": 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isNil(); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGt_isSlice(t *testing.T) {
	tests := []struct {
		name string
		c    Gt
		want bool
	}{
		{"false", Gt{"a": 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isSlice(); got != tt.want {
				t.Errorf("isSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGt_toSQL(t *testing.T) {
	type args struct {
		p *placeholders
	}
	tests := []struct {
		name    string
		c       Gt
		args    args
		want    string
		wantErr bool
	}{
		{"simple", Gt{"a": 1}, args{newPlaceholders()}, `"a" > $1`, false},
		{"simple string", Gt{"a": "1"}, args{newPlaceholders()}, `"a" > $1`, false},
		{"advanced", Gt{"a": 1, "b": 2}, args{newPlaceholders()}, `"a" > $1 AND "b" > $2`, false},
		{"in", Gt{"a": In[int]{1, 2, 3}}, args{newPlaceholders()}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.toSQL(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("toSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGte_isNil(t *testing.T) {
	tests := []struct {
		name string
		c    Gte
		want bool
	}{
		{"true", Gte{}, true},
		{"false", Gte{"a": 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isNil(); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGte_isSlice(t *testing.T) {
	tests := []struct {
		name string
		c    Gte
		want bool
	}{
		{"false", Gte{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isSlice(); got != tt.want {
				t.Errorf("isSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGte_toSQL(t *testing.T) {
	type args struct {
		p *placeholders
	}
	tests := []struct {
		name    string
		c       Gte
		args    args
		want    string
		wantErr bool
	}{
		{"simple", Gte{"a": 1}, args{newPlaceholders()}, `"a" >= $1`, false},
		{"simple string", Gte{"a": "1"}, args{newPlaceholders()}, `"a" >= $1`, false},
		{"advanced", Gte{"a": 1, "b": 2}, args{newPlaceholders()}, `"a" >= $1 AND "b" >= $2`, false},
		{"in", Gte{"a": In[int]{1, 2, 3}}, args{newPlaceholders()}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.toSQL(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("toSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIn_isNil(t *testing.T) {
	type testCase[T any] struct {
		name string
		c    In[T]
		want bool
	}
	tests := []testCase[int]{
		{"true", In[int]{}, true},
		{"false", In[int]{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isNil(); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIn_isSlice(t *testing.T) {
	type testCase[T any] struct {
		name string
		c    In[T]
		want bool
	}
	tests := []testCase[int]{
		{"true", In[int]{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isSlice(); got != tt.want {
				t.Errorf("isSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIn_toSQL(t *testing.T) {
	type args struct {
		p *placeholders
	}
	type testCase[T any] struct {
		name    string
		c       In[T]
		args    args
		want    string
		wantErr bool
	}
	tests := []testCase[int]{
		{"one", In[int]{1}, args{newPlaceholders()}, "$1", false},
		{"many", In[int]{1, 2, 3}, args{newPlaceholders()}, "$1, $2, $3", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.toSQL(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("toSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLt_isNil(t *testing.T) {
	tests := []struct {
		name string
		c    Lt
		want bool
	}{
		{"true", Lt{}, true},
		{"false", Lt{"a": 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isNil(); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLt_isSlice(t *testing.T) {
	tests := []struct {
		name string
		c    Lt
		want bool
	}{
		{"false", Lt{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isSlice(); got != tt.want {
				t.Errorf("isSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLt_toSQL(t *testing.T) {
	type args struct {
		p *placeholders
	}
	tests := []struct {
		name    string
		c       Lt
		args    args
		want    string
		wantErr bool
	}{
		{"simple", Lt{"a": 1}, args{newPlaceholders()}, `"a" < $1`, false},
		{"simple string", Lt{"a": "1"}, args{newPlaceholders()}, `"a" < $1`, false},
		{"advanced", Lt{"a": 1, "b": 2}, args{newPlaceholders()}, `"a" < $1 AND "b" < $2`, false},
		{"in", Lt{"a": In[int]{1, 2, 3}}, args{newPlaceholders()}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.toSQL(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("toSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLte_isNil(t *testing.T) {
	tests := []struct {
		name string
		c    Lte
		want bool
	}{
		{"true", Lte{}, true},
		{"false", Lte{"a": 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isNil(); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLte_isSlice(t *testing.T) {
	tests := []struct {
		name string
		c    Lte
		want bool
	}{
		{"false", Lte{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isSlice(); got != tt.want {
				t.Errorf("isSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLte_toSQL(t *testing.T) {
	type args struct {
		p *placeholders
	}
	tests := []struct {
		name    string
		c       Lte
		args    args
		want    string
		wantErr bool
	}{
		{"simple", Lte{"a": 1}, args{newPlaceholders()}, `"a" <= $1`, false},
		{"simple string", Lte{"a": "1"}, args{newPlaceholders()}, `"a" <= $1`, false},
		{"advanced", Lte{"a": 1, "b": 2}, args{newPlaceholders()}, `"a" <= $1 AND "b" <= $2`, false},
		{"in", Lte{"a": In[int]{1, 2, 3}}, args{newPlaceholders()}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.toSQL(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("toSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNq_isNil(t *testing.T) {
	tests := []struct {
		name string
		c    Nq
		want bool
	}{
		{"true", Nq{}, true},
		{"false", Nq{"a": 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isNil(); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNq_isSlice(t *testing.T) {
	tests := []struct {
		name string
		c    Nq
		want bool
	}{
		{"false", Nq{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isSlice(); got != tt.want {
				t.Errorf("isSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNq_toSQL(t *testing.T) {
	type args struct {
		p *placeholders
	}
	tests := []struct {
		name    string
		c       Nq
		args    args
		want    string
		wantErr bool
	}{
		{"simple", Nq{"a": 1}, args{newPlaceholders()}, `"a" <> $1`, false},
		{"simple string", Nq{"a": "1"}, args{newPlaceholders()}, `"a" <> $1`, false},
		{"advanced", Nq{"a": 1, "b": 2}, args{newPlaceholders()}, `"a" <> $1 AND "b" <> $2`, false},
		{"in", Nq{"a": In[int]{1, 2, 3}}, args{newPlaceholders()}, `"a" NOT IN ($1, $2, $3)`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.toSQL(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("toSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOr_isNil(t *testing.T) {
	tests := []struct {
		name string
		c    Or
		want bool
	}{
		{"true", Or{}, true},
		{"false", Or{Eq{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isNil(); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOr_isSlice(t *testing.T) {
	tests := []struct {
		name string
		c    Or
		want bool
	}{
		{"false", Or{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isSlice(); got != tt.want {
				t.Errorf("isSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOr_toSQL(t *testing.T) {
	type args struct {
		p *placeholders
	}
	tests := []struct {
		name    string
		c       Or
		args    args
		want    string
		wantErr bool
	}{
		{"simple", Or{Eq{"a": 1}}, args{newPlaceholders()}, `("a" = $1)`, false},
		{"simple", Or{Eq{"a": 1}, Eq{"b": 2}}, args{newPlaceholders()}, `("a" = $1) OR ("b" = $2)`, false},
		{"advanced", Or{Eq{"a": 1, "b": 2}, Eq{"c": 3, "d": 4}}, args{newPlaceholders()}, `("a" = $1 AND "b" = $2) OR ("c" = $3 AND "d" = $4)`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.toSQL(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("toSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toAndOrSQL(t *testing.T) {
	type args struct {
		c []Condition
		o string
		p *placeholders
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"and", args{[]Condition{Eq{"a": 1}, Eq{"b": 2}}, " AND ", newPlaceholders()}, `("a" = $1) AND ("b" = $2)`, false},
		{"or", args{[]Condition{Eq{"a": 1}, Eq{"b": 2}}, " OR ", newPlaceholders()}, `("a" = $1) OR ("b" = $2)`, false},
		{"error", args{[]Condition{new(badCondition)}, " AND ", newPlaceholders()}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toAndOrSQL(tt.args.c, tt.args.o, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("toAndOrSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toAndOrSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toEqNqSQL(t *testing.T) {
	type args struct {
		c map[string]any
		o bool
		p *placeholders
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"eq", args{map[string]any{"a": 1}, true, newPlaceholders()}, `"a" = $1`, false},
		{"nq", args{map[string]any{"a": 1}, false, newPlaceholders()}, `"a" <> $1`, false},
		{"bad slice", args{map[string]any{"a": &badSliceCondition{}}, false, newPlaceholders()}, "", true},
		{"good condition", args{map[string]any{"a": Eq{"a": 1}}, false, newPlaceholders()}, `"a" = $1`, false},
		{"bad condition", args{map[string]any{"a": &badCondition{}}, false, newPlaceholders()}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toEqNqSQL(tt.args.c, tt.args.o, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("toEqNqSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toEqNqSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toLtGtSQL(t *testing.T) {
	type args struct {
		c map[string]any
		o string
		p *placeholders
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"", args{map[string]any{"a": 1}, ">", newPlaceholders()}, `"a" > $1`, false},
		{"", args{map[string]any{"a": 1}, "<", newPlaceholders()}, `"a" < $1`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toLtGtSQL(tt.args.c, tt.args.o, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("toLtGtSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toLtGtSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type badCondition struct{}

func (b *badCondition) isNil() bool {
	return false
}

func (b *badCondition) isSlice() bool {
	return false
}

func (b *badCondition) toSQL(_ *placeholders) (string, error) {
	return "", errors.New("error")
}

type badSliceCondition struct{}

func (b *badSliceCondition) isNil() bool {
	return false
}

func (b *badSliceCondition) isSlice() bool {
	return true
}

func (b *badSliceCondition) toSQL(_ *placeholders) (string, error) {
	return "", errors.New("error")
}
