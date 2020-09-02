package shell

// SortInt will sort a list using shell sort algorithm
//
// Depends on gap sequence worst case: O(N^2).
// For the gap sequence (2^k + 1) we get O(N^{2/3}),
// the implemented sequence is the3-smooth gap sequence which we get
// O(N (logN)²) which is the best known case.

func SortInt(list []int) {
	n := len(list)

	// Create gaps using
	// gaps := []int{1}
	// k := 1
	// for {
	// 	interval := power(2, k) + 1
	// 	if interval > n-1 {
	// 		break
	// 	}
	// 	gaps = append([]int{interval}, gaps...)
	// 	k++
	// }
	gaps := Psmooth(3, []int{1, n})

	for _, interval := range gaps {
		for i := interval; i < n; i += interval {
			j := i
			for j > 0 {
				if list[j-interval] > list[j] {
					list[j-interval], list[j] = list[j], list[j-interval]
				}
				j -= interval
			}
		}
	}
}

func power(exponent, index int) int {
	power := 1
	for index > 0 {
		if index&1 != 0 {
			power *= exponent
		}
		index >>= 1
		exponent *= exponent
	}
	return power
}

// PrimeFactors Get all prime factors of a given number n
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func max(slice []int) int {
	var max int
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return max
}

// Psmooth returns p-smooth numbers on a given interval [a,b]
func Psmooth(p int, interval []int) []int {
	psmooth := []int{}
	for i := interval[0]; i <= interval[1]; i++ {
		if max(PrimeFactors(i)) <= p {
			psmooth = append(psmooth, i)
		}
	}
	return psmooth
}
