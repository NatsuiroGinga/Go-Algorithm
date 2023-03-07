package search

import (
	"fmt"
	"strings"
)

func generateLogin() {
	var s1, s2 string
	ret := "zzzzzzzzzz"

	_, _ = fmt.Scan(&s1, &s2)

	for i := range s1 {
		for j := range s2 {
			tmp := s1[:i+1] + s2[:j+1]
			if strings.Compare(tmp, ret) == 1 {
				tmp = ret
			}
		}
	}

	fmt.Println(ret)
}
