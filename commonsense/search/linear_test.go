package search_test

import (
	"testing"

	"github.com/stevedesilva/dsa-golang.git/commonsense/search"
	"github.com/stretchr/testify/assert"
)

func TestLinearShouldReturnNilIfItemNotFound(t *testing.T) {
	_, err := search.SearchUnsorted("Hi", []search.Any{})
	assert.Equal(t, err, search.NotFound)
}
