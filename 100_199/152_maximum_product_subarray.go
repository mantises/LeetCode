package _100_199

/*
Given an integer array nums, find the contiguous subarray within an array
(containing at least one number) which has the largest product.

Example 1:
Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Example 2:
Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

*/

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minTmp, maxTmp := nums[0], nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= 0 {
			minTmp, maxTmp = min(minTmp*nums[i], nums[i]), max(maxTmp*nums[i], nums[i])
		} else if nums[i] < 0 {
			minTmp, maxTmp = min(maxTmp*nums[i], nums[i]), max(minTmp*nums[i], nums[i])
		}
		res = max(res, maxTmp)
	}
	return res
}

/*
1.子问题
2.转移方程
3.初始化、边界值
4.求解
*/
