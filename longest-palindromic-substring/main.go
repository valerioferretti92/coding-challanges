package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome2("azsesq"))
}

/***************** Recursive solution *****************/

func longestPalindromes1(s string) string {
	// Base case 1
	if isPalindrome(string(s)) {
		return s
	}

	// Divide
	lpal := longestPalindrome2(s[1:])
	rpal := longestPalindrome2(s[0 : len(s)-1])

	// Conquer
	if len(lpal) > len(rpal) {
		return lpal
	} else {
		return rpal
	}
}

/***************** Linear solution *****************/

func longestPalindrome2(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := max(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/***************** Quadratic solution *****************/

func longestPalindrome3(s string) string {
	if len(s) <= 1 {
		return s
	}

	for i := len(s); i > 0; i-- {
		palindrome, ok := lookUpPalindromes(i, s)
		if ok {
			return palindrome
		}
	}

	return ""
}

func lookUpPalindromes(l int, s string) (string, bool) {
	var start, end int = 0, l

	for end <= len(s) {
		if isPalindrome(s[start:end]) {
			return s[start:end], true
		}

		start++
		end++
	}

	return "", false
}

/***************** Utils *****************/

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
