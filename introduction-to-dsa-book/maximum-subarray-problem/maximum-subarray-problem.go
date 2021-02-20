package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const uintsize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

const (
	maxint  = 1<<(uintsize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	minint  = -maxint - 1         // -1 << 31 or -1 << 63
	maxuint = 1<<uintsize - 1     // 1<<32 - 1 or 1<<64 - 1
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("[ERROR]: one and only one string of space separated prices must be provided")
	}

	prices := parseCmdLine(strings.Split(strings.Trim(os.Args[1], " \n\t"), " "))
	increases := make([]int, 0, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		increases = append(increases, int(prices[i]-prices[i-1]))
	}

	left, right, netValue := maximumSubarray(increases, 0, len(increases)-1)
	fmt.Printf("[SOLUTION]: BUY on day %d, sell on day %d, net value = %d\n", left, right+1, netValue)
}

func parseCmdLine(args []string) []uint {
	prices := make([]uint, 0, len(args)-1)
	for _, arg := range args[1:] {
		price, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("%s is not a number!\n", arg)
		}
		if price < 0 {
			log.Fatalf("%d is not a positive number!\n", price)
		}
		prices = append(prices, uint(price))
	}
	return prices
}

func maximumSubarray(myArray []int, left, right int) (maxLeft, maxRight int, sum int) {
	// Base case
	if left == right {
		return left, right, myArray[left]
	}

	// Divide
	mid := int(math.Floor(float64(left+right) / float64(2)))

	// Conquer
	leftLow, leftHigh, leftSum := maximumSubarray(myArray, left, mid)
	rigthLow, rightHigh, rigthSum := maximumSubarray(myArray, mid+1, right)
	crossLow, crossHigh, crossSum := maxCrossingSubarray(myArray, left, right, mid)

	// Combine
	if leftSum >= crossSum && leftSum >= rigthSum {
		return leftLow, leftHigh, leftSum
	}
	if rigthSum >= leftSum && rigthSum >= crossSum {
		return rigthLow, rightHigh, rigthSum
	}
	return crossLow, crossHigh, crossSum
}

func maxCrossingSubarray(myArray []int, left, right, mid int) (maxLeft, maxRight int, sum int) {
	leftMaxSum := minint
	leftSum := 0
	maxLeft = mid
	for i := mid; i >= left; i-- {
		leftSum = leftSum + myArray[i]
		if leftSum > leftMaxSum {
			leftMaxSum = leftSum
			maxLeft = i
		}
	}
	rightMaxSum := minint
	rightSum := 0
	maxRight = mid + 1
	for j := mid + 1; j <= right; j++ {
		rightSum = rightSum + myArray[j]
		if rightSum > rightMaxSum {
			rightMaxSum = rightSum
			maxRight = j
		}
	}

	return maxLeft, maxRight, leftMaxSum + rightMaxSum
}
