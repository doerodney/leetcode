package main

import (
	"fmt"
)

type ListNode struct {
  Val int
  Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Len int
}


func (l * LinkedList) insert(val int) {
	node := &ListNode{Val: val, Next: nil}
	if l.Len == 0 {
		l.Head = node
	} else {
		ptr := l.Head
		for i := 0; i < l.Len; i++ {
			if ptr.Next == nil {
				ptr.Next = node
				break
			} else {
				ptr = ptr.Next
			}
		}
	}
	l.Len++
}


func getNumber(list *ListNode) int {
	multiplier := 1
	number := 0
	ptr := list
	for ptr != nil {
		number += ptr.Val * multiplier
		ptr = ptr.Next
		multiplier *= 10
	}
	return number
}


func main() {
	xl := LinkedList{}
	yl := LinkedList{}

	xd := []int{2, 4, 3,}
	yd := []int{5, 6, 4,}

	for _, x := range(xd) {
		xl.insert(x)
	}

	xn := getNumber(xl.Head)

	for _, y := range(yd) {
		yl.insert(y)
	}
	yn := getNumber(yl.Head)


	fmt.Printf("Slice: %v, Number: %d\n", xd, xn)
	fmt.Printf("Slice: %v, Number: %d\n", yd, yn)
}