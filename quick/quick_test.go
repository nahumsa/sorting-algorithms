package quick

import (
	"testing"

	"github.com/nahumsa/sorting-algorithms/sorttest"
)

func TestInsertionSortInt(t *testing.T) {
	sorttest.TestInt(t, SortInt)
}

func BenchmarkInsertionSortInt(b *testing.B) {
	sorttest.BenchmarkInt(b, SortInt)
}
