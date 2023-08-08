package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t1 := TreeNode{Val: 1}
	t2 := TreeNode{Val: 2}
	t3 := TreeNode{Val: 3}
	t4 := TreeNode{Val: 4}
	t5 := TreeNode{Val: 5}

	t1.Left = &t2
	t2.Right = &t3
	t3.Left = &t4
	t1.Right = &t5

	fmt.Printf("Maximum depth is %d\n", maxDepth(&t1))
}

func maxDepth(root *TreeNode) int {
	return recursiveMaxDepth(root, 0)
}

func recursiveMaxDepth(node *TreeNode, value int) int {
	if node == nil {
		return value
	}

	value = value + 1
	lvalue := recursiveMaxDepth(node.Left, value)
	rvalue := recursiveMaxDepth(node.Right, value)

	if lvalue > value {
		value = lvalue
	}
	if rvalue > value {
		value = rvalue
	}
	return value
}
