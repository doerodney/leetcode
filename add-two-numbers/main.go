package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
  Val int
  Next *ListNode
}

func insert(list *ListNode, val int) {
	// If Next points to itself, this is the head of the list.
	if list.Next == list {
		list.Val = val
		list.Next = nil
	} else {
		// Iterate the nodes until Next is nil:
		ptr := list
		for (ptr.Next != nil) {
			ptr = ptr.Next
		}
		// Found the end of the list.  Add the node:
		ptr.Next = &ListNode{val, nil}
	}
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

func newLinkedList() *ListNode {
	p := &ListNode{}
	// Mark this as the head of the list (Next points to itself)
	p.Next = p
	return p
}


// Given 807, return [7, 0, 8]
func getDigits(num int) []int {
	var s []int
	asTxt := strconv.FormatInt(int64(num), 10)
	chars := strings.Split(asTxt, "")
	nChars := len(chars)
	for i := (nChars - 1); i >= 0; i-- {
		digit, _ := strconv.Atoi(chars[i])
		s = append(s, digit)
	}

	return s
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
}


func main() {
	xl := newLinkedList()
	yl := newLinkedList()

	xd := []int{2, 4, 3,}
	yd := []int{5, 6, 4,}

	for _, x := range(xd) {
		insert(xl, x)
	}

	xn := getNumber(xl)

	for _, y := range(yd) {
		insert(yl, y)
	}
	yn := getNumber(yl)

	sum := xn + yn
	digits := getDigits(sum)
	
	fmt.Printf("Slice: %v, Number: %d\n", xd, xn)
	fmt.Printf("Slice: %v, Number: %d\n", yd, yn)
	fmt.Printf("Slice: %v, Number: %d\n", digits, sum)
}