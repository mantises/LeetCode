package _500_599

import (
	"math"
)

/*
Given a matrix consists of 0 and 1, find the distance of the nearest 0 for each cell.

The distance between two adjacent cells is 1.

Example 1:
Input:
[[0,0,0],
 [0,1,0],
 [0,0,0]]

Output:
[[0,0,0],
 [0,1,0],
 [0,0,0]]
Example 2:

Input:
[[0,0,0],
 [0,1,0],
 [1,1,1]]

Output:
[[0,0,0],
 [0,1,0],
 [1,2,1]]

Note:
The number of elements of the given matrix will not exceed 10,000.
There are at least one 0 in the given matrix.
The cells are adjacent in only four directions: up, down, left and right.
*/
func updateMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	res := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		t := make([]int, 0, n)
		for j := 0; j < n; j++ {
			v := math.MaxInt32
			if matrix[i][j] == 0 {
				v = 0
			}
			t = append(t, v)
		}
		res = append(res, t)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i >= 1 {
				res[i][j] = min(res[i][j], res[i-1][j]+1)
			}
			if j >= 1 {
				res[i][j] = min(res[i][j], res[i][j-1]+1)
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i < m-1 {
				res[i][j] = min(res[i][j], res[i+1][j]+1)
			}
			if j < n-1 {
				res[i][j] = min(res[i][j], res[i][j+1]+1)
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
