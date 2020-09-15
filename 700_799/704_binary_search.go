package _700_799

/*

Given a sorted (in ascending order) integer array nums of n elements and a target
value, write a function to search target in nums. If target exists, then return
its index, otherwise return -1.

*/

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			low = mid - 1
		} else if nums[mid] < target {
			high = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
