package search

import "errors"

var (
	NotFound = errors.New("Not Found")
)

type Any interface{}

func SearchUnsorted(Value Any, array []Any) (int, error) {

	return 0, NotFound
}
