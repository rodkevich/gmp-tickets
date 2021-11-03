package user

import (
	"github.com/rodkevich/gmp-tickets/lib/filter"
	"net/url"
)

type Filter struct {
	Base filter.Common

	Field1 string
	Field2 string
	Field3 string
}

func Filters(queries url.Values) *Filter {
	f := filter.New(queries)
	return &Filter{
		Base:   *f,
		Field1: queries.Get("1"),
		Field2: queries.Get("2"),
		Field3: queries.Get("3"),
	}
}
