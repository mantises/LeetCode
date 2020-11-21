package _01_99

import "math"

/*
Given an integer array nums, find the contiguous subarray (containing at least one number)
which has the largest sum and return its sum.

Follow up: If you have figured out the O(n) solution, try coding another solution using the
divide and conquer approach, which is more subtle.


Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:
Input: nums = [1]
Output: 1

Example 3:
Input: nums = [0]
Output: 0

Example 4:
Input: nums = [-1]
Output: -1

Example 5:
Input: nums = [-2147483647]
Output: -2147483647

Constraints:
1 <= nums.length <= 2 * 10^4
-2^31 <= nums[i] <= 2^31 - 1

*/

func maxSubArrayByDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pre, maxSum := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		pre = int(math.Max(float64(pre+nums[i]), float64(nums[i])))
		maxSum = int(math.Max(float64(maxSum), float64(pre)))
	}
	return maxSum
}

func maxSubArrayByDivideAndConquer(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}
	if len(nums) == 1 {
		return nums[0]
	}
	mid := int(len(nums) / 2)
	left := maxSubArrayByDivideAndConquer(nums[0:mid])
	right := maxSubArrayByDivideAndConquer(nums[mid:len(nums)])
	crossSum := getCrossSum(nums)
	sum := int(math.Max(math.Max(float64(crossSum), float64(left)), float64(right)))
	return sum
}

func getCrossSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := math.MinInt32, math.MinInt32
	mid := int(len(nums) / 2)
	sum := 0
	for i := mid - 1; i >= 0; i-- {
		sum += nums[i]
		left = int(math.Max(float64(sum), float64(left)))
	}
	sum = 0
	for i := mid; i < len(nums); i++ {
		sum += nums[i]
		right = int(math.Max(float64(sum), float64(right)))
	}
	return left + right
}
