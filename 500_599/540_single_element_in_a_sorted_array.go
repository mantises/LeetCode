package _500_599

/*
You are given a sorted array consisting of only integers where every element appears
exactly twice, except for one element which appears exactly once. Find this single
element that appears only once.

Follow up: Your solution should run in O(log n) time and O(1) space.


Example 1:
Input: nums = [1,1,2,3,3,4,4,8,8]
Output: 2

Example 2:
Input: nums = [3,3,7,7,10,11,11]
Output: 10


Constraints:
1 <= nums.length <= 10^5
0 <= nums[i] <= 10^5

*/

func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if mid == low || mid == high {
			return nums[mid]
		}
		if nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1] {
			return nums[mid]
		} else if (mid%2 == 0 && nums[mid] == nums[mid+1]) ||
			(mid%2 == 1 && nums[mid] == nums[mid-1]) {
			low = mid + mid%2
		} else {
			high = mid - mid%2
		}
	}
	return 0
}

func singleNonDuplicateByXOR(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}
