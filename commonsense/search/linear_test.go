package search_test

import (
	"testing"

	"github.com/stevedesilva/dsa-golang.git/commonsense/search"
	"github.com/stretchr/testify/assert"
)

func TestLinearShouldReturnNilIfItemNotFoundEmptyArray(t *testing.T) {
	var value search.Any = "Hi"
	_, err := search.UnsortedSearch(&value, []search.Any{})
	assert.Equal(t, search.NotFound, err)
}

func TestLinearShouldReturnIndexWhenItemNotFound(t *testing.T) {
	_, err := search.UnsortedSearch("Missing", []search.Any{"Qs", "Foo", "bar"})
	assert.Equal(t, search.NotFound, err)
}

func TestLinearShouldReturnIndexWhenItemFound(t *testing.T) {
	res, err := search.UnsortedSearch("Hi", []search.Any{"Qs", "Foo", "bar", "Hi"})
	assert.Nil(t, err)
	assert.Equal(t, 3, res)
}
