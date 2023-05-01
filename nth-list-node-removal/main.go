package main

func main() {
	n4 := &ListNode{
		Val:  5,
		Next: nil,
	}
	n3 := &ListNode{
		Val:  4,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  3,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  2,
		Next: n2,
	}
	n0 := &ListNode{
		Val:  1,
		Next: n1,
	}

	removeNthFromEnd(n0, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	nodes := make([]*ListNode, 0)
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	target := nodes[len(nodes)-n]

	// Removing first element of the list
	if target == nodes[0] {
		return target.Next
	}

	// Removing last element of the list
	if target == nodes[len(nodes)-1] {
		nodes[len(nodes)-2].Next = nil
		return head
	}

	left := nodes[len(nodes)-n-1]
	right := nodes[len(nodes)-n+1]
	left.Next = right
	return head
}
