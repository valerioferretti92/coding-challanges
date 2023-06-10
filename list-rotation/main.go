package main

import "fmt"

func main() {
	l1 := ListNode{Val: 0}
	l2 := ListNode{Val: 1}
	l3 := ListNode{Val: 2}
	l4 := ListNode{Val: 3}

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4

	head := rotateRight(&l1, 2)

	for e := head; e != nil; e = e.Next {
		fmt.Println(e.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	// If head is nil or the rotation factor is 0, return head
	if head == nil || k == 0 {
		return head
	}

	// Computing length and tail
	length := 1
	var tail *ListNode = nil
	for curr := head.Next; curr != nil; curr = curr.Next {
		length++

		if curr.Next == nil {
			tail = curr
		}
	}

	// Computing rotation factor
	k = k % length

	// If the correct rotation factor is 0, return head
	if k == 0 {
		return head
	}

	// Computing the new list head
	var newHead *ListNode = head
	var newTail *ListNode = nil
	for i := 0; i < length-k; i++ {
		newTail = newHead
		newHead = newHead.Next
	}

	// Connecting with left hand side of the list
	newTail.Next = nil
	tail.Next = head
	return newHead
}
