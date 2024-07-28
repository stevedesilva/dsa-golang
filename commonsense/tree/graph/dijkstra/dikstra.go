package dijkstra

import "github.com/stevedesilva/dsa-golang.git/commonsense/tree/heap/min/comparator"

type Dijkstra struct {
}

func (d *Dijkstra) ShortestPath(startCity, destinationCity *City) []string {
	cheapestPriceFromStartingCity := make(map[string]int)
	cheapestPreviousStopoverCity := make(map[string]string)

	visitedCities := make(map[string]City)
	unvisitedCities := comparator.New[City](func(i, j City) bool {
		return cheapestPriceFromStartingCity[i.Name] < cheapestPriceFromStartingCity[j.Name]
	})

	// initialize
	cheapestPriceFromStartingCity[startCity.Name] = 0
	currentCity := startCity

	// for each route in starting
	for currentCity != nil {
		// process current city
		visitedCities[currentCity.Name] = *currentCity

		for adjacentCity, adjacentCityCost := range currentCity.Routes {
			if _, ok := visitedCities[adjacentCity.Name]; !ok {
				err := unvisitedCities.Insert(*adjacentCity)
				if err != nil {
					return nil
				}
			}
			// cost from current city to adjacentCity
			priceThroughStartingCity := cheapestPriceFromStartingCity[currentCity.Name] + adjacentCityCost
			if existingCheapestPriceForCity, ok := cheapestPriceFromStartingCity[adjacentCity.Name]; !ok || priceThroughStartingCity < existingCheapestPriceForCity {
				cheapestPriceFromStartingCity[adjacentCity.Name] = priceThroughStartingCity
				cheapestPreviousStopoverCity[adjacentCity.Name] = currentCity.Name
			}
		}

		// chose next city
		city, err := unvisitedCities.Delete()
		if err != nil {
			break
		}
		currentCity = city
	}

	// reverse  cheapestPreviousStopoverCity and current
	result := make([]string, 0, len(cheapestPreviousStopoverCity))
	currentCityName := destinationCity.Name
	for currentCityName != startCity.Name {
		result = append(result, currentCityName)
		currentCityName = cheapestPreviousStopoverCity[currentCityName]
	}
	result = append(result, startCity.Name)
	// reverse result
	for i := 0; i < len(result)/2; i++ {
		j := len(result) - i - 1
		result[i], result[j] = result[j], result[i]
	}

	return result
}
