package insertion

// SortInt will sort a list using insertion sort algorithm
//
// O(N^2) without binary serach and O(N log N) with binary search.
func SortInt(list []int) {
	var sorted []int

	for _, value := range list {
		sorted = insert(sorted, value)
	}

	// Copying to the list
	for i, val := range sorted {
		list[i] = val
	}

}

// insert will add a value into a sorted list if this value is less than
// an element of the list will add it to the left and move all list right
// otherwise, will put the value last.

func insert(sorted []int, value int) []int {
	for i, sortedValue := range sorted {
		if value < sortedValue {
			return append(sorted[:i], append([]int{value}, sorted[i:]...)...)
		}
	}
	return append(sorted, value)
}
