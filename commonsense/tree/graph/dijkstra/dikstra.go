package dijkstra

type Dijkstra struct {
}

func (d *Dijkstra) ShortestPath(start, destination *City) []string {
	// setup - previousCity, visitedCities, unvisitedCities, cheapestPriceFromStartingCity
	previousVisitedCity := make(map[string]string)
	visitedCities := make(map[string]bool)
	unvisitedCities := make([]string, 0, len(start.Routes))
	currentCity := start
	// for each route in starting
	// calculate shortest

	// reverse  previousVisitedCity and current
	return []string{"A", "B", "C"}
}
