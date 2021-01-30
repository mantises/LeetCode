package _500_599

/*
Given two words word1 and word2, find the minimum number of steps required to make
word1 and word2 the same, where in each step you can delete one character in either
string.

Example 1:
Input: "sea", "eat"
Output: 2
Explanation: You need one step to make "sea" to "ea" and another step to make "eat"
to "ea".

Note:
The length of given words won't exceed 500.
Characters in given words can only be lower-case letters.

*/

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, 0)
	for i := 0; i <= len(word1); i++ {
		dp = append(dp, make([]int, len(word2)+1))
	}
	for i := 0; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = i + j
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
