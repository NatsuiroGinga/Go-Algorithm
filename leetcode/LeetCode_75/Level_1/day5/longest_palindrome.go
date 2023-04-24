package day5

func longestPalindrome(s string) int {
	var (
		size    = len(s)
		charMap = make(map[byte]bool, size) // use map to store the char
		count   = 0                         // count the number of repeated char
	)

	for _, char := range []byte(s) { // convert string to byte slice
		if _, ok := charMap[char]; ok { // if the char is in the map, delete it and count++
			count++
			delete(charMap, char)
		} else { // if the char is not in the map, add it to the map
			charMap[char] = true
		}
	}

	if size >= count*2+1 { // if the length of the string is greater than the number of repeated char * 2 + 1
		return count*2 + 1 // return count*2 + 1
	}

	return count * 2
}
