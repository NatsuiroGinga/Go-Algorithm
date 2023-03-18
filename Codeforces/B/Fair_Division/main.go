package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		t, n uint8
		//candies []int
	)
	reader := bufio.NewReader(os.Stdin)
	//writer := bufio.NewWriter(os.Stdout)
	t, _ = reader.ReadByte()

	for t -= '0'; t > 0; t-- {
		fmt.Scan(&n)
	}
}
