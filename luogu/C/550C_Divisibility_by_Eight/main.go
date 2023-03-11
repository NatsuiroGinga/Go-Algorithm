package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Scanln(&n)

	for i := 0; i < len(n); i++ {
		c1 := n[i]
		if c1 == '0' || c1 == '8' {
			fmt.Println("YES")
			fmt.Print(string(c1))
			return
		}

		for j := i + 1; j < len(n); j++ {
			var c2 = n[j]
			if ((c1-'0')*10+(c2-'0'))%8 == 0 {
				fmt.Println("YES")
				fmt.Print(string(c1) + string(c2))
				return
			}

			for k := j + 1; k < len(n); k++ {
				c3 := n[k]
				if ((c1-'0')*100+(c2-'0')*10+(c3-'0'))%8 == 0 {
					fmt.Println("YES")
					fmt.Print(string(c1) + string(c2) + string(c3))
					return
				}
			}

		}
	}

	fmt.Print("NO")
}
