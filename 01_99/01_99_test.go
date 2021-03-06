package _01_99

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{-1, -2, -3, -4, -5}
	target := -8
	t.Log(twoSumByDoublePointer(nums, target))
}

func TestRemoveNthFromEnd(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	t.Log(removeNthFromEndByDoublePointer(l1, 1))
}

func TestMyAtoi(t *testing.T) {
	t.Log(myAtoi("+3"))
}

func TestSearchInRotatedSortedArray(t *testing.T) {
	nums := []int{3, 1}
	target := 1
	t.Log(search(nums, target))
}

func TestFindPivot(t *testing.T) {
	nums := []int{2, 3, 5, 10, -2, -1, 0, 1}
	t.Log(findPivot(nums))
	t.Log(searchByFindPivot(nums, 4))
}

func TestSearchRotatedSortedArrayII(t *testing.T) {
	nums := []int{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
	}
	t.Log(searchByFindPivotII(nums, 2))
}

func TestThreeSum(t *testing.T) {
	nums := []int{-2, 0, 1, 1, 2}
	t.Log(threeSum(nums))
}

func TestMergeSortedArray(t *testing.T) {
	nums1 := []int{1, 2, 3, 5, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	nums2 := []int{2, 4, 5, 6, 7, 9, 10}
	merge(nums1, 5, nums2, 7)
}

func TestMergeKSortedLists(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 4}
	l3 := &ListNode{Val: 5}
	l1.Next = l2
	l2.Next = l3

	l4 := &ListNode{Val: 2}
	l5 := &ListNode{Val: 13}
	l4.Next = l5

	l6 := &ListNode{Val: 9}
	l7 := &ListNode{Val: 16}
	l6.Next = l7
	lists := []*ListNode{l1, l4, l6}
	s := mergeKLists(lists)
	fmt.Println(s)
}

func TestSortColors(t *testing.T) {
	nums := []int{2, 2, 2, 1, 0, 0, 1, 2, 0, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func TestMaximumSubarray(t *testing.T) {
	nums := []int{1, 2, -1}
	fmt.Println(maxSubArrayByDP(nums))
	fmt.Println(maxSubArrayByDivideAndConquer(nums))
}

func TestEditDistance(t *testing.T) {
	word1 := "intention"
	word2 := "execution"
	fmt.Println(minDistanceV2(word1, word2))
}

func TestLongestPalindromicSubstring(t *testing.T) {
	s := "ababaabbcc"
	fmt.Println(longestPalindrome(s))
}

func TestMergeIntervals(t *testing.T) {
	intervals := [][]int{
		{2, 3}, {1, 9}, {10, 12}, {10, 12},
	}
	fmt.Println(mergeInterval(intervals))
}

func TestCanJump(t *testing.T) {
	fmt.Println(canJump([]int{2, 2, 0, 1}))
}

func TestJump(t *testing.T) {
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
}

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("ABCDBEDFFBGHDI", "BDG"))
}

func TestRotateList(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 4}
	l3 := &ListNode{Val: 5}
	l4 := &ListNode{Val: 9}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	fmt.Println(rotateRight(l1, 2))
}

func TestMySqrt(t *testing.T) {
	fmt.Println(mySqrt(8))
}

func TestDeleteDuplicatesII(t *testing.T) {
	l1 := &ListNode{Val: 4}
	l2 := &ListNode{Val: 4}
	l3 := &ListNode{Val: 4}
	l4 := &ListNode{Val: 4}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	fmt.Println(deleteDuplicatesII(l1))
}

func TestReverseBetween(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	fmt.Println(reverseBetween(l1, 5, 5))
}

func TestPartition(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 2}
	l5 := &ListNode{Val: 5}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	fmt.Println(partition(l1, 3))
}

func TestFirstMissingPositive(t *testing.T) {
	fmt.Println(firstMissingPositive([]int{2, 1, 4, 3, 0, 4}))
}

func TestRecoverTree(t *testing.T) {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  0,
				Left: nil,
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		Right: &TreeNode{
			Val: 9999,
			Left: &TreeNode{
				Val: 101,
			},
		},
	}
	recoverTree(root)
	fmt.Println(root)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("wpwwpwaxwdq"))
}
