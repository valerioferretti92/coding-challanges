package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("p"))
}

func lengthOfLongestSubstring(s string) int {
	longest := uint(0)
	currentLength := uint(0)
	charMap := make(map[rune]uint)

	for idx := 0; idx < len(s); {
		currRune, size := utf8.DecodeRuneInString(s[idx:])

		storedIdx, ok := charMap[currRune]
		if !ok {
			charMap[currRune] = uint(idx)
			currentLength++
			idx += size
			continue
		}

		if currentLength >= longest {
			longest = currentLength
		}

		charMap = make(map[rune]uint)
		idx = int(storedIdx)
		currentLength = 0
		idx += size
	}

	if currentLength >= longest {
		longest = currentLength
	}
	return int(longest)
}
