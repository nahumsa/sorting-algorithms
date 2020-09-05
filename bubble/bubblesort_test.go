package bubble

import (
	"testing"

	"github.com/nahumsa/sorting-algorithms/sorttest"
)

func TestBubbleSortInt(t *testing.T) {
	sorttest.TestInt(t, SortInt)
}

func BenchmarkBubbleSortInt(b *testing.B) {
	sorttest.BenchmarkInt(b, SortInt)
}

// func TestBubbleSortString(t *testing.T) {
// 	sorttest.TestString(t, BubbleSortString)
// }

// func TestBubbleSortInterface(t *testing.T) {
// 	sorttest.TestInterface(t, BubbleSort)
// }
