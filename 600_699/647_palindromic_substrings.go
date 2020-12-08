package _600_699

/*
Given a string, your task is to count how many palindromic substrings in this string.

The substrings with different start indexes or end indexes are counted as different substrings
even they consist of same characters.

Example 1:
Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".

Example 2:
Input: "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

Note:
The input string length won't exceed 1000.

*/
func countSubstrings(s string) int {
	if len(s) <= 1 {
		return 1
	}
	count := 0
	var dp [][]bool
	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j < 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				count++
			}
		}
	}
	return count
}
