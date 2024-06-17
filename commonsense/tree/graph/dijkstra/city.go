package dijkstra

//type City struct {
//	Name   string
//	Routes map[string]Route
//}
//
//type Route struct {
//	City City
//	Cost int
//}
//
//func NewCity(Name string) *City {
//	return &City{
//		Name:   Name,
//		Routes: make(map[string]Route),
//	}
//}
//
//func (c *City) AddRoute(Name string, value int) {
//	c.Routes[Name] = Route{
//		City: City{},
//		Cost: value,
//	}
//}

type City struct {
	Name   string
	Routes map[*City]int
}

func NewCity(name string) *City {
	return &City{
		Name:   name,
		Routes: make(map[*City]int),
	}
}

func (c *City) AddRoute(city *City, price int) {
	c.Routes[city] = price
}

func (c *City) GetName() string {
	return c.Name
}

func (c *City) GetRoutes() map[*City]int {
	return c.Routes
}
