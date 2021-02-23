package _01_99

/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you
beforehand.

(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).

You are given a target value to search. If found in the array return true, otherwise
return false.

Example 1:
Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true

Example 2:
Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false

Follow up:
This is a follow up problem to Search in Rotated Sorted Array, where nums may contain
duplicates.
Would this affect the run-time complexity? How and why?
*/

//
func searchByForce(nums []int, target int) bool {
	for _, v := range nums {
		if v == target {
			return true
		}
	}
	return false
}

func searchRotatedSortedArrayII(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return true
		} else if nums[mid] == nums[low] {
			low++
		} else if nums[mid] == nums[high] {
			high--
		} else if nums[low] < nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { // right
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}

func searchByFindPivotII(nums []int, target int) bool {
	return findRecursive(nums, target, 0, len(nums)-1) != -1
}

func findRecursive(nums []int, target, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)>>1
	if nums[mid] == target {
		return mid
	}
	res := findRecursive(nums, target, low, mid-1)
	if res == -1 {
		res = findRecursive(nums, target, mid+1, high)
	}
	return res
}
