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

func ReportNode(node *TreeNode) {
	if node != nil {
		fmt.Printf("Node: %d\n", node.Val)
		if node.Left != nil {
			fmt.Printf("\tNode.Left.Val: %d\n", node.Left.Val)
		} else {
			fmt.Printf("\tNode.Left.Val: nil\n")
		}
		if node.Right != nil {
			fmt.Printf("\tNode.Right.Val: %d\n", node.Right.Val)
		} else {
			fmt.Printf("\tNode.Right.Val: nil\n")
		}
	}
}


func TraverseBST(root *TreeNode) {
	var stack stack.Stack
	current := root

	for current != nil || !stack.IsEmpty() {
		// Push all the left nodes into the stack:
		for current != nil {
			stack.Push(current)
			current = current.Left
		}

		// Pop the left node stack and visit its right nodes:
		if current == nil && !stack.IsEmpty() {
			obj, _ := stack.Pop()
			// Dynamic cast type from interface{} to *TreeNode:
			current = obj.(*TreeNode)
			ReportNode(current)
			current = current.Right
		}
	}
}

func main() {
	values := []int{4,2,1,3,6,5,7}

	var root *TreeNode
	for _, value := range values {
		root = insertIntoBST(root, value)
	}

	TraverseBST(root)	
}