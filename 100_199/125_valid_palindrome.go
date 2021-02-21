package _100_199

import "strings"

/*
Given a string s, determine if it is a palindrome, considering only alphanumeric
characters and ignoring cases.


Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.


Constraints:
1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.

*/
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if !isValid(s[left]) {
			left++
			continue
		}
		if !isValid(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return left >= right
}

func isValid(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
