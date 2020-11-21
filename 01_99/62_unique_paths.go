package _01_99

/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is trying to reach
the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?
Example 1:
Input: m = 3, n = 7
Output: 28

Example 2:
Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down

Example 3:
Input: m = 7, n = 3
Output: 28

Example 4:
Input: m = 3, n = 3
Output: 6

*/

/*
时间复杂度：o(n^2)
空间复杂度：o(n^2)
*/
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	count := make([][]int, n)
	for i := range count {
		count[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		count[0][i] = 1
	}
	for i := 0; i < n; i++ {
		count[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			count[j][i] = count[j-1][i] + count[j][i-1]
		}
	}
	return count[n-1][m-1]
}

/*
时间复杂度：o(n^2)
空间复杂度：o(n)
*/
func uniquePathsV2(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	count := make([]int, n)
	for i := 0; i < n; i++ {
		count[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			count[j] += count[j-1]
		}
	}
	return count[n-1]
}
