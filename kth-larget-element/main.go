package main

import (
	"container/heap"
	"fmt"
)

func main() {
	values := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(values, 4))
}

type sorted_array []int

func (s sorted_array) Len() int {
	return len(s)
}

func (s sorted_array) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s sorted_array) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *sorted_array) Push(value any) {
	*s = append((*s), value.(int))
}

func (s *sorted_array) Pop() any {
	value := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return value
}

func findKthLargest(nums []int, k int) int {
	var sa sorted_array = make([]int, 0)
	heap.Init(&sa)

	for _, num := range nums {
		heap.Push(&sa, num)
	}

	value := 0
	for i := 0; i < k; i++ {
		value = heap.Pop(&sa).(int)
	}

	return value
}
