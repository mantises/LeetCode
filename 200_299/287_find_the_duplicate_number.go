package _200_299

/*
Given an array of integers nums containing n + 1 integers where each integer is in the range
[1, n] inclusive.

There is only one duplicate number in nums, return this duplicate number.

Follow-ups:

How can we prove that at least one duplicate number must exist in nums?
Can you solve the problem without modifying the array nums?
Can you solve the problem using only constant, O(1) extra space?
Can you solve the problem with runtime complexity less than O(n2)?

Example 1:

Input: nums = [1,3,4,2,2]
Output: 2
Example 2:

Input: nums = [3,1,3,4,2]
Output: 3
Example 3:

Input: nums = [1,1]
Output: 1
Example 4:

Input: nums = [1,1,2]
Output: 1

Constraints:
	2 <= n <= 3 * 104
	nums.length == n + 1
	1 <= nums[i] <= n
	All the integers in nums appear only once except for precisely one integer which appears two or more times.

*/

func findDuplicateByForce(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return -1
}

func findDuplicateByBinarySearch(nums []int) int {
	low, high := 1, len(nums)-1
	res := -1
	for low <= high {
		mid := low + (high-low)>>1
		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		if count <= mid {
			low = mid + 1
		} else {
			high = mid - 1
			res = mid
		}
	}
	return res
}

func findDuplicateByDoublePointers(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
