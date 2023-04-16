package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(maxArea(array[:]))
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		result = max(result, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
