package main

var input = bufio.NewReader(os.Stdin)
var output = bufio.NewWriter(os.Stdout)

var n int

func main() {
	defer output.Flush()
	fmt.Fscan(input, &n)
}
