package main

import "log"

func reverseStr(s string, k int) string {
	b := []byte(s)

	for i := 0; i < len(b); i += 2 * k {
		if i+k-1 < len(b) {
			reverse(b, i, i+k-1)
			continue
		}
		reverse(b, i, len(b)-1)
	}

	return string(b)
}

func reverse(s []byte, low, high int) {
	for ; low < high; low, high = low+1, high-1 {
		s[low], s[high] = s[high], s[low]
	}
}

func main() {
	s := "abcdefg"
	k := 4
	str := reverseStr(s, k)
	log.Println(str)
}
