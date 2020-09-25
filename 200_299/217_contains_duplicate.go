package _200_299

import "sort"

/*
Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array,
and it should return false if every element is distinct.

Example 1:
Input: [1,2,3,1]
Output: true

Example 2:
Input: [1,2,3,4]
Output: false

Example 3:
Input: [1,1,1,3,3,4,3,2,4,2]
Output: true

*/

// sort and compare
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// use hash map
func containsDuplicateByHashMap(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = true
		}
	}
	return false
}
