package _200_299

/*
Given an m x n binary matrix filled with 0's and 1's, find the largest square containing
only 1's and return its area.


Example 1:
Input: matrix = [
				 ["1","0","1","0","0"],
				 ["1","0","1","1","1"],
				 ["1","1","1","1","1"],
				 ["1","0","0","1","0"]]
Output: 4

Example 2:
Input: matrix = [["0","1"],["1","0"]]
Output: 1

Example 3:
Input: matrix = [["0"]]
Output: 0

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] is '0' or '1'.

*/
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	res := make([][]int, 0)
	for i := 0; i <= m; i++ {
		res = append(res, make([]int, n+1))
	}
	maxSide := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				res[i][j] = min(res[i-1][j-1], min(res[i-1][j], res[i][j-1])) + 1
			}
			maxSide = max(maxSide, res[i][j])
		}
	}
	return maxSide * maxSide
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
