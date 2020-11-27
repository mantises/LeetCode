package _100_199

import "math"

/*
You are a professional robber planning to rob houses along a street. Each house has a certain
amount of money stashed, the only constraint stopping you from robbing each of them is that
adjacent houses have security system connected and it will automatically contact the police
if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house,
determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.

Example 2:
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.


Constraints:
0 <= nums.length <= 100
0 <= nums[i] <= 400

*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	money := make([]int, len(nums))
	money[0], money[1] = nums[0], int(math.Max(float64(nums[0]), float64(nums[1])))
	for i := 2; i < len(nums); i++ {
		money[i] = int(math.Max(float64(money[i-1]), float64(money[i-2]+nums[i])))
	}
	return money[len(nums)-1]
}

func robV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	first, second := nums[0], int(math.Max(float64(nums[0]), float64(nums[1])))
	for i := 2; i < len(nums); i++ {
		first, second = second, int(math.Max(float64(first+nums[i]), float64(second)))
	}
	return second
}
