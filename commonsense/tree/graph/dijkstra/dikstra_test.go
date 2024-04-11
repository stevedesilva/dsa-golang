package dijkstra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {

	// Given setup city with routes
	atlanta := NewCity("Atlanta")
	chicago := NewCity("Chicago")
	boston := NewCity("Boston")
	denver := NewCity("Denver")
	elPaso := NewCity("El Paso")

	atlanta.AddRoute(boston.Name, 100)
	atlanta.AddRoute(denver.Name, 160)
	boston.AddRoute(chicago.Name, 120)
	boston.AddRoute(denver.Name, 180)
	chicago.AddRoute(elPaso.Name, 80)
	denver.AddRoute(chicago.Name, 40)
	denver.AddRoute(elPaso.Name, 140)
	elPaso.AddRoute(boston.Name, 100)

	// When
	d := Dijkstra{}
	shortestPath := d.ShortestPath(atlanta, elPaso)
	// Then
	expected := []string{"Atlanta", "Denver", "Chicago", "El Paso"}
	assert.Equal(t, expected, shortestPath)
}
