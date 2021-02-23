package _500_599

import (
	"fmt"
	"testing"
)

func Test01Matrix(t *testing.T) {
	matrix := [][]int{
		{0, 0, 1, 1},
		{0, 1, 1, 1},
		{0, 0, 1, 1},
	}
	updateMatrix(matrix)
}

func TestMinDistance(t *testing.T) {
	fmt.Println(minDistance("eat", "ates"))
}

func TestFindLongestWord(t *testing.T) {
	fmt.Println(findLongestWord("adpldaplet",
		[]string{"app", "adapt", "apple", "apldat"}))
}

func TestSingleNonDuplicate(t *testing.T) {
	fmt.Println(singleNonDuplicate([]int{5, 7, 7, 8, 8}))
}
