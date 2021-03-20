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

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 78,
				Left: &TreeNode{
					Val:  1443546,
					Left: nil,
					Right: &TreeNode{
						Val: 10,
					},
				},
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 9999,
			Left: &TreeNode{
				Val: 119,
			},
		},
	}
	fmt.Println(diameterOfBinaryTree(root))
}

func TestFindMode(t *testing.T) {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 78,
				Left: &TreeNode{
					Val:  1443546,
					Left: nil,
					Right: &TreeNode{
						Val: 10,
					},
				},
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 9999,
			Left: &TreeNode{
				Val: 119,
			},
		},
	}
	fmt.Println(findMode(root))
}
