package dijkstra

import "github.com/stevedesilva/dsa-golang.git/commonsense/tree/heap/min/comparator"

type Dijkstra struct {
}

func (d *Dijkstra) ShortestPath(startCity, destinationCity *City) []string {
	cheapestPriceFromStartingCity := make(map[string]int)
	cheapestPreviousStopoverCity := make(map[string]string)

	visitedCities := make(map[string]City)
	unvisitedCities := comparator.New[City](func(i, j City) bool {
		return cheapestPriceFromStartingCity[i] < cheapestPriceFromStartingCity[j]
	})

	// initialize
	cheapestPriceFromStartingCity[startCity.Name] = 0
	currentCity := startCity

	// for each route in starting
	for currentCity != nil {
		// process current city
		visitedCities[currentCity.Name] = *currentCity

		for adjacentCityName, adjacentCityRoute := range currentCity.Routes {
			if _, ok := visitedCities[adjacentCityName]; !ok {
				err := unvisitedCities.Insert(adjacentCityRoute.City)
				if err != nil {
					return nil
				}
				//unvisitedCities = append(unvisitedCities, adjacentCityName)

			}
			var adjacentCityCost int = adjacentCityRoute.Cost

			// cost from current city to adjacentCity
			priceThroughStartingCity := cheapestPriceFromStartingCity[currentCity.Name] + adjacentCityCost
			if existingCheapestPriceForCity, ok := cheapestPriceFromStartingCity[adjacentCityName]; !ok || priceThroughStartingCity < existingCheapestPriceForCity {
				cheapestPriceFromStartingCity[adjacentCityName] = priceThroughStartingCity
				cheapestPreviousStopoverCity[adjacentCityName] = currentCity.Name
			}
		}

		// chose next city

		city, err := unvisitedCities.Delete()
		if err != nil {
			return nil
		}
		currentCity = city
	}
	// calculate shortest

	// reverse  cheapestPreviousStopoverCity and current
	return []string{"A", "B", "C"}
}

//func removeElement(nameToRemove string, cities []*City) []*City {
//	for i, v := range cities {
//		if v.Name == nameToRemove {
//			return append(cities[:i], cities[i+1:]...)
//		}
//	}
//	return cities
//}
