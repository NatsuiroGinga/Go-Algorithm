package main

import "fmt"

func main() {
	fmt.Println(32 << (^uint(0) >> 63))
}
