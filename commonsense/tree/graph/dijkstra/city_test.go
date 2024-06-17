package dijkstra

import "testing"

func TestCity_AddRoute(t *testing.T) {
	c := NewCity("City")
	o := NewCity("City_2")
	c.AddRoute(o, 0)
	//if c.Routes["City"] != 0 {
	//	t.Errorf("City.AddRoute() = %v, want %v", c.Routes["City"], 0)
	//}
}
