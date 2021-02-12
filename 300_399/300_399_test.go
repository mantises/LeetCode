package _300_399

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{2, 3, 5}
	fmt.Println(coinChange(coins, 15))
}

func TestLengthOfLIS(t *testing.T) {
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(nums))
}

func TestMaxCoins(t *testing.T) {
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums))
}

func TestMaxProfit(t *testing.T) {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}

func TestHouseRobberIII(t *testing.T) {
	root := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Val: 4,
			},
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
		Val: 2,
	}
	fmt.Println(rob(root))
}

func TestIntegerBreak(t *testing.T) {
	fmt.Println(integerBreak(10))
}

func TestRemoveDuplicate(t *testing.T) {
	fmt.Println(removeDuplicateLetters("svcavbcxs"))
}

func TestIsSubsequence(t *testing.T) {
	fmt.Println(isSubsequence("asx", "axstvx"))
}

func TestCreateMaximumNumber(t *testing.T) {
	fmt.Println(maxNumber([]int{5, 6, 9, 8, 7}, []int{8, 9, 4, 6}, 2))
}
