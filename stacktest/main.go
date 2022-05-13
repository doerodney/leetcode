package main

import (
	"fmt"
	"github.com/doerodney/leetcode/stack"
)


func main() {
	var int_stack stack.Stack
	int_stack.Push(0)
	int_stack.Push(1)
	int_stack.Push(2)

	for !int_stack.IsEmpty() {
		value, ok := int_stack.Pop()
		if ok {
			n := value.(int)
			fmt.Println(n)
		}
	}

	var str_stack stack.Stack
	str_stack.Push("zero")
	str_stack.Push("one")
	str_stack.Push("two")

	for !str_stack.IsEmpty() {
		value, ok := str_stack.Pop()
		if ok {
			s := value.(string)
			fmt.Println(s)
		}
	}
}
