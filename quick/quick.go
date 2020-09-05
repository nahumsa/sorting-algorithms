package quick

// SortInt will sort a list using quick sort algorithm
//
// O(N^2).
func SortInt(list []int) {

	quickSorter(list, 0, len(list)-1)

}

func quickSorter(elements []int, below int, upper int) {
	if below < upper {
		var part int
		part = divideParts(elements, below, upper)
		quickSorter(elements, below, part-1)
		quickSorter(elements, part+1, upper)
	}
}

func divideParts(elements []int, below int, upper int) int {

	center := elements[upper]

	i := below
	for j := below; j < upper; j++ {
		if elements[j] <= center {
			elements[i], elements[j] = elements[j], elements[i]
			i++
		}
	}

	elements[i], elements[upper] = elements[upper], elements[i]
	return i
}
