package main

func main() {

}

func findMatrix(nums []int) [][]int {
	freq := make(map[int]int)
	var ans [][]int

	for _, num := range nums {
		freq[num]++
	}

	for len(freq) > 0 {
		var line []int

		for k, v := range freq {
			if v > 0 {
				line = append(line, k)
				freq[k]--
			} else {
				delete(freq, k)
			}
		}

		if len(line) > 0 {
			ans = append(ans, line)
		}
	}

	return ans
}
