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
//func NewCity(name string) *City {
//	return &City{
//		Name:   name,
//		Routes: make(map[string]Route),
//	}
//}
//
//func (c *City) AddRoute(name string, value int) {
//	c.Routes[name] = Route{
//		City: City{},
//		Cost: value,
//	}
//}

type City struct {
	name   string
	routes map[*City]int
}

func NewCity(name string) *City {
	return &City{
		name:   name,
		routes: make(map[*City]int),
	}
}

func (c *City) AddRoute(city *City, price int) {
	c.routes[city] = price
}

func (c *City) GetName() string {
	return c.name
}

func (c *City) GetRoutes() map[*City]int {
	return c.routes
}
