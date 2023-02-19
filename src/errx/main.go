package errx

import (
	"fmt"
)

type Error struct {
	Err string
	Parent []any
	Values []any
}

func New(err string) Error {
	e := Error{Err: err}
	return e
}

func (e Error) P(p any) Error {
	e.Parent = []any{p}
	return e
}

func (e Error) V(v ...any) Error {
	e.Values = v
	return e
}

func (e Error) Error() string {
	var(
		vals []any
		ret string
	)

	if len(e.Parent) > 0 {
		ret = fmt.Sprintf("%s: ", e.Parent[0])
	}

	vals = append(vals, e.Values...)

	return ret + fmt.Sprintf(
		e.Err,
		vals...,
	)
}

func (e Error) String() string {
	return e.Error()
}

