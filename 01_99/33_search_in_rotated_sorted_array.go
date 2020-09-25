package _01_99

/*
You are given an integer array nums sorted in ascending order, and an integer target.

Suppose that nums is rotated at some pivot unknown to you beforehand
(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

If target is found in the array return its index, otherwise, return -1.

*/

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		}
		// which side is in ascending order
		if nums[mid] >= nums[low] { // left
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { //right
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func searchByFindPivot(nums []int, target int) int {
	pivot := findPivot(nums)
	high := len(nums) - 1
	if nums[pivot] <= target && target <= nums[high] {
		index := BinarySearch(nums[pivot:], target)
		if index == -1 {
			return -1
		} else {
			return pivot + index
		}
	} else {
		index := BinarySearch(nums[:pivot], target)
		if index == -1 {
			return -1
		} else {
			return index
		}
	}
}

func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func findPivot(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
