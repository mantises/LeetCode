package _300_399

/*
Given a string s, remove duplicate letters so that every letter appears once and only once. You
must make sure your result is the smallest in lexicographical order among all possible results.

Note: This question is the same as 1081:
https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/

Example 1:
Input: s = "bcabc"
Output: "abc"

Example 2:
Input: s = "cbacdcbc"
Output: "acdb"


Constraints:
1 <= s.length <= 10^4
s consists of lowercase English letters.

*/
func removeDuplicateLetters(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}
	var stack []byte
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}
