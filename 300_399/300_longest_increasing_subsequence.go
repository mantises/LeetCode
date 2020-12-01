package _300_399

/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.
A subsequence is a sequence that can be derived from an array by deleting some or no elements
without changing the order of the remaining elements. For example, [3,6,2,7] is a subsequence
of the array [0,3,1,6,2,2,7].

Example 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1


Constraints:
1 <= nums.length <= 2500
-10^4 <= nums[i] <= 10^4

Follow up:
Could you come up with the O(n^2) solution?
Could you improve it to O(n*log(n)) time complexity?
*/
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = 1
	}
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				res[i] = max(res[i], res[j]+1)
			}
		}
		maxLength = max(maxLength, res[i])
	}
	return maxLength
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
