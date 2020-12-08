package _300_399

/*
Given n balloons, indexed from 0 to n-1. Each balloon is painted with a number on it
represented by array nums. You are asked to burst all the balloons. If the you burst
balloon i you will get nums[left] * nums[i] * nums[right] coins. Here left and right
are adjacent indices of i. After the burst, the left and right then becomes adjacent.

Find the maximum coins you can collect by bursting the balloons wisely.

Note:
You may imagine nums[-1] = nums[n] = 1. They are not real therefore you can not burst them.
0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100.

Example:
Input: [3,1,5,8]
Output: 167
Explanation: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
             coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
*/
func maxCoins(nums []int) int {
	tmp := append(append([]int{1}, nums...), 1)
	dp := make([][]int, len(tmp))
	for i := range dp {
		dp[i] = make([]int, len(tmp))
	}
	for i := 3; i <= len(tmp); i++ {
		for j := 0; j <= len(tmp)-i; j++ {
			for k := j + 1; k < j+i-1; k++ {
				left := dp[j][k]
				right := dp[k][j+i-1]
				dp[j][j+i-1] = max(dp[j][j+i-1], left+tmp[j]*tmp[k]*tmp[j+i-1]+right)
			}
		}
	}
	return 0
}
