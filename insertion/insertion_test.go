package insertion

import (
	"testing"

	"github.com/nahumsa/sorting-algorithms/sorttest"
)

func TestInsertionSortInt(t *testing.T) {
	sorttest.TestInt(t, SortInt)
}

func TestInsertionSortIntBS(t *testing.T) {
	sorttest.TestInt(t, SortIntBS)
}
func BenchmarkInsertionSortInt(b *testing.B) {
	sorttest.BenchmarkInt(b, SortInt)
}

func BenchmarkInsertionSortIntBT(b *testing.B) {
	sorttest.BenchmarkInt(b, SortIntBS)
}

// func TestBubbleSortString(t *testing.T) {
// 	sorttest.TestString(t, BubbleSortString)
// }

// func TestBubbleSortInterface(t *testing.T) {
// 	sorttest.TestInterface(t, BubbleSort)
// }
