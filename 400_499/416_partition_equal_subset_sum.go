package _400_499

/*
Given a non-empty array nums containing only positive integers, find if the array can be
partitioned into two subsets such that the sum of elements in both subsets is equal.

Example 1:
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

Example 2:
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.

Constraints:
1 <= nums.length <= 200
1 <= nums[i] <= 100
*/

func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 1; i <= len(nums); i++ {
		for j := target; j >= nums[i-1]; j-- {
			dp[j] = dp[j] || dp[j-nums[i-1]]
		}
	}
	return dp[target]
}
