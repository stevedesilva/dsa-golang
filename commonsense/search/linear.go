package search

func UnsortedSearch(Value Any, array []Any) (int, error) {
	for idx, item := range array {
		if found := Value == item; found {
			return idx, nil
		}
	}
	return 0, NotFound
}
