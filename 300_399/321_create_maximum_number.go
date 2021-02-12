package _300_399

/*
Given two arrays of length m and n with digits 0-9 representing two numbers. Create the maximum
number of length k <= m + n from digits of the two. The relative order of the digits from the
same array must be preserved. Return an array of the k digits.

Note: You should try to optimize your time and space complexity.

Example 1:
Input:
nums1 = [3, 4, 6, 5]
nums2 = [9, 1, 2, 5, 8, 3]
k = 5
Output:
[9, 8, 6, 5, 3]

Example 2:
Input:
nums1 = [6, 7]
nums2 = [6, 0, 4]
k = 5
Output:
[6, 7, 6, 0, 4]

Example 3:
Input:
nums1 = [3, 9]
nums2 = [8, 9]
k = 3
Output:
[9, 8, 9]

*/
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)
	res := make([]int, 0)
	for i := max(0, k-n); i <= min(m, k); i++ {
		s1 := getMaxStock(nums1, i)
		s2 := getMaxStock(nums2, k-i)
		merged := mergeStocks(s1, s2)
		if lexicographicalLess(res, merged) {
			res = merged
		}
	}
	return res
}

func lexicographicalLess(nums1, nums2 []int) bool {
	for i := 0; i < len(nums1) && i < len(nums2); i++ {
		if nums1[i] != nums2[i] {
			return nums1[i] < nums2[i]
		}
	}
	return len(nums1) < len(nums2)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func getMaxStock(nums []int, length int) []int {
	if length == 0 {
		return []int{}
	}
	dropNum := len(nums) - length
	stack := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for dropNum > 0 && len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			stack = stack[:len(stack)-1]
			dropNum--
		}
		if len(stack) < length {
			stack = append(stack, nums[i])
		} else {
			dropNum--
		}
	}
	return stack
}

func mergeStocks(nums1, nums2 []int) []int {
	res := make([]int, len(nums1)+len(nums2))
	for i := 0; i < len(res); i++ {
		if lexicographicalLess(nums1, nums2) {
			res[i], nums2 = nums2[0], nums2[1:]
		} else {
			res[i], nums1 = nums1[0], nums1[1:]
		}
	}
	return res
}
