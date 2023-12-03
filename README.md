# pqr
Another postgres ORM like package. Deeply inspired by [reform](https://github.com/go-reform/reform) and [squirrel](https://github.com/Masterminds/squirrel).

[![Tests](https://github.com/skamenetskiy/pqr/actions/workflows/tests.yml/badge.svg)](https://github.com/skamenetskiy/pqr/actions/workflows/tests.yml)

## Installation
```bash
go get -u github.com/skamenetskiy/pqr
```

## Sample

```go
package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/skamenetskiy/pqr"
)

// Job repository definition.
type Job struct {
	ID         int64      `db:"id,pk"`
	StartedAt  *time.Time `db:"started_at"`
	FinishedAt *time.Time `db:"finished_at"`
	Errored    bool       `db:"errored"`
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	jobs := pqr.New[Job, int64]("jobs",
		pqr.WithLogger(pqr.DefaultLogger),
		pqr.WithStandardSQL(db),
	)
	if err != nil {
		panic(err)
	}

	// find a single job with id = 11
	job, err := jobs.FindOne(11)
	if err != nil {
		panic(err)
	}

	// update operation
	job.Errored = true
	if err = jobs.Update(pqr.UpdateParams[Job]{
		Condition: pqr.Eq{"id": 11},
		Columns:   []string{"errored"},
		Items:     []*Job{job},
	}); err != nil {
		panic(err)
	}

	// find multiple elements
	jl, err := jobs.Find(pqr.FindParams{
		Condition: pqr.Eq{
			"id": pqr.In[int64]{11, 5, 6},
		},
	})
	fmt.Println(jl)

	// run in transaction
	err = jobs.Transaction(func(tx pqr.Transaction[Job, int64]) error {
		c, e := tx.Count()
		fmt.Println("count", c, e)
		return e
	})
	if err != nil {
		panic(err)
	}
}

```

## How to avoid runtime reflection
Very simple! Your repository struct must implement `Valueable` and `Pointable` interfaces. That's it!