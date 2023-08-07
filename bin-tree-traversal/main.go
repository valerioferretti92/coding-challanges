package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t1 := TreeNode{Val: 30}
	t2 := TreeNode{Val: 20}
	t3 := TreeNode{Val: 40}
	t4 := TreeNode{Val: 10}
	t5 := TreeNode{Val: 50}

	t1.Left = &t2
	t1.Right = &t3
	t2.Left = &t4
	t3.Right = &t5

	fmt.Printf("%v\n", inorderTraversal(&t1))
}

func inorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)
	return recursiveTraversal(values, root)
}

func recursiveTraversal(values []int, node *TreeNode) []int {
	if node == nil {
		return values
	}

	values = recursiveTraversal(values, node.Left)
	values = append(values, node.Val)
	values = recursiveTraversal(values, node.Right)
	return values
}

func iterativeTraversal(values []int, curr *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	visited := make(map[*TreeNode]bool)

	for curr != nil {
		if curr.Left != nil && !visited[curr.Left] {
			stack = append(stack, curr)
			curr = curr.Left
			continue
		}

		values = append(values, curr.Val)
		visited[curr] = true

		if curr.Right != nil {
			curr = curr.Right
			continue
		}

		if len(stack) == 0 {
			curr = nil
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	return values
}
