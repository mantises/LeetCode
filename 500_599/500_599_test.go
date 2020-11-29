package _500_599

import "testing"

func Test01Matrix(t *testing.T) {
	matrix := [][]int{
		{0, 0, 1, 1},
		{0, 1, 1, 1},
		{0, 0, 1, 1},
	}
	updateMatrix(matrix)
}
