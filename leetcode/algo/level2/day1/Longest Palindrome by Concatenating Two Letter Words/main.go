package main

func main() {

}

func longestPalindrome(words []string) int {
	fre := make([][]int, 26)
	var ans int
	for i := range fre {
		fre[i] = make([]int, 26)
	}

	for _, w := range words {
		if fre[w[1]-'b'][w[0]-'a'] > 0 {
			ans += 4
			fre[w[1]-'b'][w[0]-'a']--
		} else {
			fre[w[0]-'a'][w[1]-'a']++
		}
	}

	for i := range fre {
		if fre[i][i] > 0 {
			ans += 2
		}
	}

	return ans
}
