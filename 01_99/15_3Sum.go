package _01_99

import "sort"

/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.

Notice that the solution set must not contain duplicate triplets.

Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Example 2:
Input: nums = []
Output: []

Example 3:
Input: nums = [0]
Output: []

Constraints:

0 <= nums.length <= 3000
-10^5 <= nums[i] <= 10^5
*/

// time complexity: o(n^2), space complexity: o(1)
// three pointer, should be careful of duplicate elements
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	if len(nums) <= 2 || nums[0] > 0 || nums[len(nums)-1] < 0 {
		return nil
	}
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 || (i > 0 && nums[i] == nums[i-1]) {
			continue
		}
		low, high := i+1, len(nums)-1
		for low < high {
			if nums[high] < 0 {
				break
			}
			if nums[i]+nums[low]+nums[high] == 0 {
				res = append(res, []int{nums[i], nums[low], nums[high]})
				low++
				high--
				for low < high && nums[low-1] == nums[low] {
					low++
				}
				for low < high && nums[high+1] == nums[high] {
					high--
				}
			} else if nums[i]+nums[low]+nums[high] < 0 {
				low++
			} else if nums[i]+nums[low]+nums[high] > 0 {
				high--
			}
		}
	}
	return res
}
