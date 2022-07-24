package search

import "errors"

var (
	NotFound = errors.New("not found")
)

type Any interface{}
