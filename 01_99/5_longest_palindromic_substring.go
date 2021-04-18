package _01_99

/*
Given a string s, return the longest palindromic substring in s.

Example 1:
Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"

Example 3:
Input: s = "a"
Output: "a"

Example 4:
Input: s = "ac"
Output: "a"


Constraints:
1 <= s.length <= 1000
s consist of only digits and English letters (lower-case and/or upper-case)
*/
func longestPalindrome(s string) string {
	var dp [][]bool
	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
	}
	maxLength := 0
	res := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j < 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				if i-j+1 > maxLength {
					res = s[j : i+1]
					maxLength = i - j + 1
				}
			}
		}
	}
	return res
}
