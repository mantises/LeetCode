package _600_699

import (
	"fmt"
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	fmt.Println(countSubstrings("aaaaa"))
}

func TestCanPlaceFlowers(t *testing.T) {
	fmt.Println(canPlaceFlowers([]int{0, 1, 0, 0}, 1))
}

func TestCheckPossibility(t *testing.T) {
	fmt.Println(checkPossibility([]int{4, 2, 1}))
}

func TestJudgeSquareSum(t *testing.T) {
	fmt.Println(judgeSquareSum(7))
}

func TestValidPalindrome(t *testing.T) {
	fmt.Println(validPalindrome("a"))
}

func TestLongestUnivaluePath(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	fmt.Println(longestUnivaluePath(root))
}

func TestTree2str(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	fmt.Println(tree2str(root))
}
