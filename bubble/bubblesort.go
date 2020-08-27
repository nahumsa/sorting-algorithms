package bubble

// SortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func SortInt(list []int) {
	for sweepNum := 0; sweepNum < len(list); sweepNum++ {
		swapped := false
		for i := 0; i < len(list)-1-sweepNum; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
