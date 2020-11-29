package _200_1299

/*
Given a m * n matrix of ones and zeros, return how many square sub-matrices have all ones.

Example 1:
Input: matrix =
[
  [0,1,1,1],
  [1,1,1,1],
  [0,1,1,1]
]
Output: 15
Explanation:
There are 10 squares of side 1.
There are 4 squares of side 2.
There is  1 square of side 3.
Total number of squares = 10 + 4 + 1 = 15.

Example 2:
Input: matrix =
[
  [1,0,1],
  [1,1,0],
  [1,1,0]
]
Output: 7
Explanation:
There are 6 squares of side 1.
There is 1 square of side 2.
Total number of squares = 6 + 1 = 7.

Constraints:
1 <= arr.length <= 300
1 <= arr[0].length <= 300
0 <= arr[i][j] <= 1

*/
func countSquares(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	res := make([][]int, 0)
	for i := 0; i <= m; i++ {
		res = append(res, make([]int, n+1))
	}
	sum := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] != 0 {
				res[i][j] = min(res[i-1][j-1], min(res[i-1][j], res[i][j-1])) + 1
			}
			sum += res[i][j]
		}
	}
	return sum
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
