package dijkstra

type City struct {
	Name   string
	Routes map[string]int
}

func NewCity(name string) *City {
	return &City{
		Name:   name,
		Routes: make(map[string]int),
	}
}

func (c *City) AddRoute(name string, value int) {
	c.Routes[name] = value
}
