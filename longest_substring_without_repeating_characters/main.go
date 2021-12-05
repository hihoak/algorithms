// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import "fmt"

func main() {
	fmt.Printf("Res is %d", lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	hashMap := make(map[uint8]int, 0)
	length, maxLength := 0, 0
	startIdxSubString := 0
	for idx := 0; idx < len(s); idx++ {
		if jdx, ok := hashMap[s[idx]]; ok && jdx >= startIdxSubString {
			if length > maxLength {
				maxLength = length
			}
			startIdxSubString = jdx + 1
			length = idx - jdx
		} else {
			length++
		}
		hashMap[s[idx]] = idx
	}
	if length > maxLength {
		maxLength = length
	}
	return maxLength
}
