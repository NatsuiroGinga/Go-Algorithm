package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type MyLinkedList struct {
	head, rear *MyListNode
	len        int
}

type MyListNode struct {
	Val  int
	Next *MyListNode
}

func Constructor() MyLinkedList {
	head := &MyListNode{Next: nil}
	return MyLinkedList{
		head: head,
		rear: head,
		len:  0,
	}
}

func (l *MyLinkedList) Get(index int) int {
	_, cur := l.find(index)

	if cur == nil {
		return -1
	}

	return cur.Val
}

/*
incr move rear to the next node and increase len

addTail means if add node to rear
*/
func (l *MyLinkedList) incr(addTail bool) {
	if addTail {
		l.rear = l.rear.Next
	}
	l.len++
}

func (l *MyLinkedList) AddAtHead(val int) {
	if l.len == 0 {
		l.AddAtTail(val)
		return
	}

	node := &MyListNode{
		Val:  val,
		Next: l.head.Next,
	}
	l.head.Next = node
	l.incr(false)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.rear.Next = &MyListNode{Val: val, Next: nil}
	l.incr(true)
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index == l.len {
		l.AddAtTail(val)
		return
	}

	pre, cur := l.find(index)
	if cur == nil {
		return
	}

	node := &MyListNode{
		Val:  val,
		Next: cur,
	}

	pre.Next = node
	l.incr(false)
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	pre, cur := l.find(index)
	if cur == nil {
		return
	}

	pre.Next = cur.Next
	l.len--

	if cur == l.rear {
		l.rear = pre
	}
}

// find the node of the index-th node in the linked list.
// If the index is invalid(index < 0 || index >= len), return pre = nil, cur = nil
func (l *MyLinkedList) find(index int) (pre, cur *MyListNode) {
	if index == l.len {
		return l.rear, nil
	}

	if index < 0 || index > l.len {
		return nil, nil
	}

	pre = l.head
	cur = l.head.Next

	for ; index > 0; index-- {
		pre = cur
		cur = cur.Next
	}

	return pre, cur
}

func (l *MyLinkedList) output() string {
	cur := l.head.Next
	var builder strings.Builder
	builder.Grow(l.len)

	for ; cur != nil; cur = cur.Next {
		builder.WriteString(strconv.Itoa(cur.Val))
		if cur != l.rear {
			builder.WriteString("->")
		}
	}

	return fmt.Sprintf("MyLinkedList : {%s}", builder.String())
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	list := Constructor()
	root := &list
	root.AddAtHead(1)
	root.AddAtTail(3)
	root.AddAtIndex(1, 2)
	log.Println("get(1) :", root.Get(1))
	root.DeleteAtIndex(1)
	log.Println("get(1) :", root.Get(1))
	log.Println(root.output())
}
