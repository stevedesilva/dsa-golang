package hashmap

import "errors"

var ErrMinimumInputRequired = errors.New("minimum input is required")
var ErrNoDuplicates = errors.New("no duplicates found in input data")

type Duplicate interface {
	FirstDuplicate() (*string, error)
}

func NewData(in []string) Duplicate {
	return &data{
		array: in,
	}
}

type data struct {
	array []string
}

func (d *data) FirstDuplicate() (*string, error) {
	if len(d.array) < 2 {
		return nil, ErrMinimumInputRequired
	}
	values := make(map[string]bool, len(d.array))
	for _, v := range d.array {
		if values[v] {
			return &v, nil
		}
		values[v] = true
	}
	return nil, ErrNoDuplicates
}

func (d *data) FirstNonDuplicate() (*string, error) {
	if len(d.array) < 2 {
		return nil, ErrMinimumInputRequired
	}
	values := make(map[string]bool, len(d.array))
	for _, v := range d.array {
		if values[v] {
			return &v, nil
		}
		values[v] = true
	}
	return nil, ErrNoDuplicates
}
