package dijkstra

type Dijkstra struct {
}

func (d *Dijkstra) ShortestPath(startCity, destinationCity *City) []string {
	cheapestPriceFromStartingCity := make(map[string]int)
	cheapestPreviousStopoverCity := make(map[string]string)

	visitedCities := make(map[string]*City)
	unvisitedCities := make([]*City, 0, 10)

	// initialize
	cheapestPriceFromStartingCity[startCity.Name] = 0
	currentCity := startCity

	// for each route in starting
	for currentCity != nil {
		// process current city
		visitedCities[currentCity.Name] = currentCity
		removeElement(currentCity.Name, unvisitedCities)

		for adjacentCityName, adjacentCityCost := range currentCity.Routes {
			if !visitedCities[adjacentCityName] {
				unvisitedCities = append(unvisitedCities, adjacentCityName)
			}
			// cost from current city to adjacentCity
			priceThroughStartingCity := cheapestPriceFromStartingCity[currentCity.Name] + adjacentCityCost
			if existingCheapestPriceForCity, ok := cheapestPriceFromStartingCity[adjacentCityName]; !ok || priceThroughStartingCity < existingCheapestPriceForCity {
				cheapestPriceFromStartingCity[adjacentCityName] = priceThroughStartingCity
				cheapestPreviousStopoverCity[adjacentCityName] = currentCity.Name
			}
		}

		// chose next city
		var city *City
		//for i,c := range unvisitedCities {
		//	cu
		//	if c.
		//}
		currentCity = city
	}
	// calculate shortest

	// reverse  cheapestPreviousStopoverCity and current
	return []string{"A", "B", "C"}
}

func removeElement(nameToRemove string, cities []*City) []*City {
	for i, v := range cities {
		if v.Name == nameToRemove {
			return append(cities[:i], cities[i+1:]...)
		}
	}
	return cities
}
