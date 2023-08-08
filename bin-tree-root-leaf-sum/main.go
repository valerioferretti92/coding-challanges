package main

import (
	"fmt"
	"math"
)

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

	fmt.Printf("Root to leaf sum is %d\n", sumNumbers(&t1))
}

func sumNumbers(root *TreeNode) int {
	sum, _ := recursiveSumNumbers(root, make([]int, 0), 0)
	return sum
}

func recursiveSumNumbers(node *TreeNode, digits []int, sum int) (int, []int) {
	if node == nil {
		return sum, digits
	}

	digits = append(digits, node.Val)

	// Leaf node
	if node.Left == nil && node.Right == nil {
		sum += buildNumber(digits)
		return sum + buildNumber(digits), digits
	}

	if node.Left != nil {
		sum, digits = recursiveSumNumbers(node.Left, digits, sum)
		digits = digits[:len(digits)-1]
	}
	if node.Right != nil {
		sum, digits = recursiveSumNumbers(node.Right, digits, sum)
		digits = digits[:len(digits)-1]
	}
	return sum, digits
}

func buildNumber(digits []int) int {
	num := float64(0)
	for i := 0; i < len(digits); i++ {
		num += math.Pow(10, float64(len(digits)-i-1)) * float64(digits[i])
	}

	return int(num)
}
