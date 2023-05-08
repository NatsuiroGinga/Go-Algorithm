package main

import (
	"log"
)

func reverseWords(s string) string {
	b := []byte(s)

	// 去空格
	slow, fast := 0, 0
	for ; fast < len(b); fast++ {
		if b[fast] != ' ' {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}

			for fast < len(b) && b[fast] != ' ' {
				b[slow] = b[fast]
				fast++
				slow++
			}
		}
	}
	b = b[:slow]

	// 整个反转
	reverse(b, 0, len(b)-1)

	// 每个单词反转
	count := 0
	i := 0
	for ; i < len(b); i++ {
		if b[i] != ' ' {
			count++
		} else {
			reverse(b, i-count, i-1)
			count = 0
		}
	}
	// 最后一个单词反转
	reverse(b, i-count, i-1)

	return string(b)
}

func reverse(s []byte, low, high int) {
	for ; low < high; low, high = low+1, high-1 {
		s[low], s[high] = s[high], s[low]
	}
}

func main() {
	s := "  hello  world  "
	words := reverseWords(s)
	log.Println(words)
}
