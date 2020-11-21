package _01_99

/*
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example 1:
Input: grid = [[1,3,1],
			   [1,5,1],
			   [4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.

Example 2:
Input: grid = [[1,2,3],
			   [4,5,6]]
Output: 12

Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100

*/

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/*
时间复杂度：o(n^2)
空间复杂度：o(n^2)
*/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	count := make([][]int, len(grid))
	for i := range count {
		count[i] = make([]int, len(grid[0]))
	}
	count[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		count[i][0] = count[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		count[0][i] = count[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			count[i][j] = grid[i][j] + min(count[i-1][j], count[i][j-1])
		}
	}
	return count[m-1][n-1]
}

/*
时间复杂度：o(n^2)
空间复杂度：o(n)
*/
func minPathSumV2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	sum := make([]int, n)
	sum[0] = grid[0][0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		sum[0] = sum[0] + grid[i][0]
		for j := 1; j < n; j++ {
			sum[j] = grid[i][j] + min(sum[j], sum[j-1])
		}
	}
	return sum[n-1]
}
