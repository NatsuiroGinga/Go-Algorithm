package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, _ := reader.ReadByte()
	t := b - '0'
	_, _, _ = reader.ReadLine()
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	flag := true

	for ; t > 0; t-- {
		s, _ := reader.ReadString('\n')
		split := strings.Split(s, " ")
		if flag {
			flag = false
		} else {
			fmt.Fprintln(writer)
		}
		_, _ = fmt.Fprintf(writer, "1 1 %s %s", split[0], split[1])
	}
}
