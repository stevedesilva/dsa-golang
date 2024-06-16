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
