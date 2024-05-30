package dijkstra

type City struct {
	Name   string
	Routes map[string]Route
}

type Route struct {
	City City
	Cost int
}

func NewCity(name string) *City {
	return &City{
		Name:   name,
		Routes: make(map[string]Route),
	}
}

func (c *City) AddRoute(name string, value int) {
	c.Routes[name] = Route{
		City: City{},
		Cost: value,
	}
}
