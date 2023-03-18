package day12

import "fmt"

func getHint(secret string, guess string) string {
	var bulls, cows int
	fre := make(map[byte]int, 10) // a map to store the frequency of each number in secret

	for i := range secret { // iterate over secret
		sChar := secret[i]
		gChar := guess[i]

		if sChar == gChar { // if the number in secret is equal to the number in guess
			bulls++
		} else {
			if fre[sChar] < 0 { // if the number in secret is less than 0, it means that the number in guess has been used
				cows++
			}
			if fre[gChar] > 0 { // if the number in guess is greater than 0, it means that the number in secret has been used
				cows++
			}
			fre[sChar]++
			fre[gChar]--
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}
