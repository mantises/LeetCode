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
	fmt.Println(maxNumber([]int{2, 5, 6, 4, 4, 0},
		[]int{7, 3, 8, 0, 6, 5, 7, 6, 2}, 15))
	// fmt.Println(lexicographicalLess([]int{9, 8, 6, 5, 3}, []int{9, 8, 4, 6, 5}, 5))
	// fmt.Println(getMaxStock([]int{4, 0, 9, 9, 0, 5, 5, 4, 7}, 4))
}

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	fmt.Println(lengthOfLongestSubstringKDistinct("anasakakad", 4))
}

func TestOddEvenList(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	printList(oddEvenList(l1))
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%v\t", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func TestPlusOne(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 9}
	l3 := &ListNode{Val: 9}
	l4 := &ListNode{Val: 9}
	l5 := &ListNode{Val: 9}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	printList(plusOne(l1))
}
