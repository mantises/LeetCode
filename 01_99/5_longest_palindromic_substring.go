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
	if len(s) <= 1 {
		return s
	}
	var dp [][]bool
	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
		dp[i][i] = true
	}
	// len
	ret := s[0:1]
	for i := 2; i <= len(s); i++ {
		// [j, j+i-1]
		for j := 0; j+i-1 < len(s); j++ {
			if s[j] != s[j+i-1] {
				dp[j][j+i-1] = false
				continue
			}
			if i == 2 && s[j] == s[j+i-1] {
				dp[j][j+i-1] = true
			} else if dp[j+1][j+i-2] == true && s[j] == s[j+i-1] {
				dp[j][j+i-1] = true
			}
			if dp[j][j+i-1] && len(ret) <= i {
				ret = s[j : j+i]
			}
		}
	}
	return ret
}
