package slices

// import github.com/jaronnie/go-generics
// generics with slices

func Filter[V any](collection []V, predicate func(V, int) bool) []V {
	var result []V

	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}
	return result
}
