package main

func main() {
}

func getNext(s string) []int {
	j := 0
	next := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}

	return next
}
