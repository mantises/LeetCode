package _600_699

/*
Given an array nums with n integers, your task is to check if it could become non-decreasing
by modifying at most one element.

We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based)
such that (0 <= i <= n - 2).

Example 1:
Input: nums = [4,2,3]
Output: true
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.

Example 2:
Input: nums = [4,2,1]
Output: false
Explanation: You can't get a non-decreasing array by modify at most one element.

Constraints:
n == nums.length
1 <= n <= 10^4
-10^5 <= nums[i] <= 10^5

*/
func checkPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	count := 0
	for i := 0; i < len(nums)-1 && count < 2; i++ {
		if nums[i] > nums[i+1] {
			count++
			if i > 0 && nums[i+1] < nums[i-1] {
				nums[i+1] = nums[i]
			} else {
				nums[i] = nums[i+1]
			}
		}
	}
	return count < 2
}
