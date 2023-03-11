package A

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	pri := make([]int, n+5)

	pri[0], pri[1] = 1, 1     // 0 and 1 are not prime
	for i := 2; i <= n; i++ { // go through all numbers from 2 to n
		if pri[i] == 0 { // i is a prime
			for j := 2; j*i <= n; j++ { // go through all multiples of i
				pri[j*i]++ // mark them as non-prime and increase the counter
			}
		}
	}

	count := 0
	// go through all numbers from 2 to n
	for _, v := range pri {
		if v == 2 { // if the number is an almost prime
			count++ // increase the counter
		}
	}

	fmt.Print(count)
}
