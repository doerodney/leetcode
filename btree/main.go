package main

import (
	"fmt"
	"github.com/doerodney/leetcode/stack"
)


 // Definition for a binary tree node.
 type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

 func insertIntoBST(root *TreeNode, val int) *TreeNode {
    node := root
    for node != nil {
		if val > node.Val {
			// Insert into the right branch:
			if node.Right == nil {
				// Insert here:
				node.Right = &TreeNode{Val:val}
				return root
			} else {
				node = node.Right 
			}
		} else {
			// Insert into the left branch:
			if node.Left == nil {
				// Insert here:
				node.Left = &TreeNode{Val:val}
				return root
			} else {
				node = node.Left
			}
		}
    }
	// Root is null.  Create a node and return it.
    return &TreeNode{Val: val}
}


func main() {
	values := []int{8, 6, 4, 2, 7}

	var root *TreeNode
	for _, value := range values {
		root = insertIntoBST(root, value)
	}

	var stack stack.Stack
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)

	for !stack.IsEmpty() {
		value, ok := stack.Pop()
		if ok {
			fmt.Println(value)
		}
	}
}