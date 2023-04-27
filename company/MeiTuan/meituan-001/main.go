package main

import (
	"fmt"
	"unicode"
)

// 小美的用户名
func main() {
	var (
		n        int
		username string
	)
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&username)
		flag := true

		if !unicode.IsLetter(rune(username[0])) {
			flag = false
		}

		if flag {
			numCount, letterCount := 0, 0

			for _, r := range username {
				if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
					flag = false
					break
				} else if unicode.IsDigit(r) {
					numCount++
				} else if unicode.IsLetter(r) {
					letterCount++
				}
			}

			if letterCount == 0 || numCount == 0 {
				flag = false
			}
		}

		if flag {
			fmt.Println("Accept")
		} else {
			fmt.Println("Wrong")
		}
	}

}
