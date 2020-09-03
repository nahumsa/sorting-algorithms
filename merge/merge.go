package merge

// SortInt will sort a list using merge sort algorithm
//
// O(N log N).
func SortInt(list []int) {

	sorted := mergeSorter(list)

	// Copying to the list
	for i, val := range sorted {
		list[i] = val
	}
}

func mergeSorter(list []int) []int {
	if len(list) < 2 {
		return list
	}
	middle := len(list) / 2

	return joinLists(mergeSorter(list[:middle]), mergeSorter(list[middle:]))
}

func joinLists(leftList, rightList []int) []int {
	num, i, j := len(leftList)+len(rightList), 0, 0
	array := make([]int, num, num)

	for k := 0; k < num; k++ {
		if i > len(leftList)-1 && j <= len(rightList)-1 {
			array[k] = rightList[j]
			j++
		} else if j > len(rightList)-1 && i <= len(leftList)-1 {
			array[k] = leftList[i]
			i++
		} else if leftList[i] < rightList[j] {
			array[k] = leftList[i]
			i++
		} else {
			array[k] = rightList[j]
			j++
		}
	}

	return array
}
