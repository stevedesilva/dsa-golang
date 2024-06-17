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

/*
private static Stream<Object[]> provideCitiesAndExpectedPaths() {
        City atlanta = new City("Atlanta");
        City boston = new City("Boston");
        City chicago = new City("Chicago");
        City denver = new City("Denver");
        City elPaso = new City("El Paso");

        atlanta.addRoutes(boston, 100);
        atlanta.addRoutes(denver, 160);
        boston.addRoutes(chicago, 120);
        boston.addRoutes(denver, 180);
        chicago.addRoutes(elPaso, 80);
        denver.addRoutes(chicago, 40);
        denver.addRoutes(elPaso, 140);
        elPaso.addRoutes(boston, 100);

        return Stream.of(
                new Object[]{atlanta, elPaso, List.of("Atlanta", "Denver", "Chicago", "El Paso")},
                new Object[]{atlanta, chicago, List.of("Atlanta", "Denver", "Chicago")},
                new Object[]{atlanta, boston, List.of("Atlanta", "Boston")},
                new Object[]{denver, elPaso, List.of("Denver", "Chicago", "El Paso")},
                new Object[]{elPaso, denver, List.of("El Paso", "Boston", "Denver")}
        );
    }
*/
// TODO create
