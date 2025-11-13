package demo03

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcabcbb")) // 3
	t.Log(lengthOfLongestSubstring("bbbbb"))    // 1
	t.Log(lengthOfLongestSubstring("pwwkew"))   // 3
	t.Log(lengthOfLongestSubstring(""))         // 0
	t.Log(lengthOfLongestSubstring(" "))        // 1
}
