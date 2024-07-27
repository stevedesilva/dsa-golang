package dijkstra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {

	// Given setup city with Routes
	atlanta := NewCity("Atlanta")
	chicago := NewCity("Chicago")
	boston := NewCity("Boston")
	denver := NewCity("Denver")
	elPaso := NewCity("El Paso")

	atlanta.AddRoute(boston, 100)
	atlanta.AddRoute(denver, 160)
	boston.AddRoute(chicago, 120)
	boston.AddRoute(denver, 180)
	chicago.AddRoute(elPaso, 80)
	denver.AddRoute(chicago, 40)
	denver.AddRoute(elPaso, 140)
	elPaso.AddRoute(boston, 100)

	// When
	d := Dijkstra{}
	shortestPath := d.ShortestPath(atlanta, elPaso)
	// Then
	expected := []string{"Atlanta", "Denver", "Chicago", "El Paso"}
	assert.Equal(t, expected, shortestPath)
}

func TestShortestPath2(t *testing.T) {

	// Given setup city with Routes
	atlanta := NewCity("Atlanta")
	chicago := NewCity("Chicago")
	boston := NewCity("Boston")
	denver := NewCity("Denver")
	elPaso := NewCity("El Paso")

	atlanta.AddRoute(boston, 100)
	atlanta.AddRoute(denver, 160)
	boston.AddRoute(chicago, 120)
	boston.AddRoute(denver, 180)
	chicago.AddRoute(elPaso, 80)
	denver.AddRoute(chicago, 40)
	denver.AddRoute(elPaso, 140)
	elPaso.AddRoute(boston, 100)

	// Whencccccbgvnutdhkfjcvnirjleuleiuhvkjddevnfercch

	d := Dijkstra{}
	shortestPath := d.ShortestPath(atlanta, denver)
	// Then
	expected := []string{"Atlanta", "Denver"}
	assert.Equal(t, expected, shortestPath)
}
