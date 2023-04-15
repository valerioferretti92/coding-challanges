package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func Test(t *testing.T) {
	var length int = 0

	length = lengthOfLongestSubstring("abcabcbb")
	if length != 3 {
		t.Fatalf("expected=3, gotten=%d", length)
	}

	length = lengthOfLongestSubstring("bbbbb")
	if length != 1 {
		t.Fatalf("expected=1, gotten=%d", length)
	}

	length = lengthOfLongestSubstring("pwwkew")
	if length != 3 {
		t.Fatalf("expected=3, gotten=%d", length)
	}

	length = lengthOfLongestSubstring(" ")
	if length != 1 {
		t.Fatalf("expected=1, gotten=%d", length)
	}

	length = lengthOfLongestSubstring("")
	if length != 0 {
		t.Fatalf("expected=0, gotten=%d", length)
	}

	var unicodestr string = "磨见見磨见見a败敗"
	fmt.Printf("len(\"磨见見磨见見a败敗\")=%d\n", len(unicodestr))
	fmt.Printf("rune_count(\"磨见見磨见見a败敗\")=%d\n", utf8.RuneCount([]byte(unicodestr)))

	length = lengthOfLongestSubstring(unicodestr)
	if length != 6 {
		t.Fatalf("expected=6, gotten=%d", length)
	}
}
