package search_test

import (
	"testing"

	"github.com/stevedesilva/dsa-golang.git/commonsense/search"
	"github.com/stretchr/testify/assert"
)

func TestLinearShouldReturnNilIfItemNotFound(t *testing.T) {
	_, err := search.UnsortedSearch("Hi", []search.Any{})
	assert.Equal(t, err, search.NotFound)
}

func TestLinearShouldReturnIndexWhenItemFound(t *testing.T) {
	valueToFind := "Hi"
	res, err := search.UnsortedSearch(valueToFind, []search.Any{"Qs", "Foo", "bar", valueToFind})
	assert.Nil(t, err)
	assert.Equal(t, res, 3)
}
