package main

import (
	"fmt"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

/******************** Cubic solution *******************/

func threeSum(nums []int) [][]int {
	triplets := make(map[string][]int)

	nums = sort(nums)
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				tripletId := buildTripletId(nums[i], nums[j], nums[k])
				if _, ok := triplets[tripletId]; !ok {
					triplets[tripletId] = make([]int, 0)
					triplets[tripletId] = append(triplets[tripletId], nums[i])
					triplets[tripletId] = append(triplets[tripletId], nums[j])
					triplets[tripletId] = append(triplets[tripletId], nums[k])
				}

				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}

	results := make([][]int, 0)
	for _, triplet := range triplets {
		results = append(results, triplet)
	}
	return results
}

// Complexity: = O(n^2)
func sort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		smallest := 100001
		index := i

		for j := i; j < len(nums); j++ {
			if nums[j] < smallest {
				smallest = nums[j]
				index = j
			}
		}

		nums[i], nums[index] = nums[index], nums[i]
	}
	return nums
}

func buildTripletId(i, j, k int) string {
	return fmt.Sprintf("%d@%d@%d", i, j, k)
}
