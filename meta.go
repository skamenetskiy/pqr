package pqr

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type meta struct {
	name             string
	key              keyInfo
	columnsWithPK    []column
	columnsWithoutPK []column
	columnsMap       map[string]column
}

type column struct {
	name  string
	typ   string
	isPK  bool
	index int
}

func parseMeta(name string, t any) (*meta, error) {
	const tagName = "db"
	r := &meta{
		name: name,
	}
	to := reflect.TypeOf(t).Elem()
	if to.Kind() != reflect.Struct {
		return nil, errors.New("element	 is not a struct")
	}
	for i := 0; i < to.NumField(); i++ {
		f := to.Field(i)
		c, e := parseTag(f.Tag.Get(tagName))
		if e != nil {
			return nil, e
		}
		c.index = i
		c.typ = f.Type.String()
		r.columnsWithPK = append(r.columnsWithPK, c)
		if !c.isPK {
			r.columnsWithoutPK = append(r.columnsWithoutPK, c)
		} else {
			r.key.index = c.index
			r.key.name = c.name
		}
	}
	if len(r.columnsWithPK) == 0 {
		return nil, fmt.Errorf("non of the struct fields had %s tag", tagName)
	}
	r.columnsMap = make(map[string]column, len(r.columnsWithPK))
	for _, c := range r.columnsWithPK {
		r.columnsMap[c.name] = c
	}
	return r, nil
}

func columnsToString(cc []column) string {
	r := make([]string, len(cc))
	for i, c := range cc {
		r[i] = `"` + c.name + `"`
	}
	return strings.Join(r, ", ")
}

func parseTag(tag string) (column, error) {
	pair := strings.Split(tag, ",")
	c := column{}
	if len(pair) == 0 {
		return c, errors.New("invalid 'db' tag value")
	}
	if len(pair) > 0 {
		c.name = pair[0]
	}
	if len(pair) > 1 && pair[1] == "pk" {
		c.isPK = true
	}
	return c, nil
}
