package day9

type MyQueue struct {
	/*
		output is used to store the output elements
		input is used to store the input elements
	*/
	output, input []int
}

func Constructor() MyQueue {
	return MyQueue{
		output: make([]int, 0, 20),
		input:  make([]int, 0, 20),
	}
}

func (q *MyQueue) Push(x int) {
	q.input = append(q.input, x)
}

func (q *MyQueue) Pop() int {
	peek := q.Peek()                      // peek the top element
	q.output = q.output[:len(q.output)-1] // pop the top element from output
	return peek
}

func (q *MyQueue) Peek() int {
	if len(q.output) == 0 { // output is empty
		// move all elements from input to output
		for len(q.input) > 0 {
			top := q.input[len(q.input)-1]     // pop the top element from input
			q.output = append(q.output, top)   // push the top element to output
			q.input = q.input[:len(q.input)-1] // pop the top element from input
		}
	}
	// return the top element of output
	return q.output[len(q.output)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.output) == 0 && len(q.input) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
