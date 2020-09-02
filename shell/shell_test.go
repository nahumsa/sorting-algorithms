package shell

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

func TestPsmooth(t *testing.T) {
	want := [][]int{
		[]int{1, 2, 4, 8, 16, 32, 64},
		[]int{9, 12, 16, 18, 24, 27},
		[]int{25, 27, 30, 32, 36, 40, 45},
		[]int{40, 42, 44, 45, 48, 49, 50, 54, 55, 56, 60, 63, 64, 66},
	}

	ps := []int{2, 3, 5, 11}

	intervals := [][]int{
		[]int{1, 70},
		[]int{9, 30},
		[]int{25, 46},
		[]int{40, 67},
	}

	for i := 0; i < len(ps); i++ {
		result := Psmooth(ps[i], intervals[i])
		errorCount := 0

		for j := 0; j < len(want[i]); j++ {
			if errorCount >= 5 {
				t.Fatalf("additional errors omitted for brevity")
			}
			if result[j] != want[i][j] {
				errorCount++
				t.Errorf("p = %d : result[%d] = %d; want %d", ps[i], j, result[j], want[i][j])
			}
		}
	}
}
