package _01_99

import "sort"

/*
Given an array of integers nums and an integer target, return indices of the
two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may
not use the same element twice.

You can return the answer in any order.
*/

// time complexity: o(n^2), space complexity: o(1)
func twoSumByForce(nums []int, target int) []int {
	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			if target-v == nums[k] {
				return []int{i, k}
			}
		}
	}
	return []int{}
}

// time complexity: o(nlogn), space complexity: o(n)
func twoSumByHashMap(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{i, k}
		}
		m[v] = i
	}
	return nil
}

// time complexity: o(nlogn), space complexity: o(1)
func twoSumByDoublePointer(nums []int, target int) []int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)
	head := 0
	tail := len(nums) - 1
	for head < tail {
		if tmp[head]+tmp[tail] < target {
			head++
		} else if tmp[head]+tmp[tail] > target {
			tail--
		} else {
			head = tmp[head]
			tail = tmp[tail]
			break
		}
	}
	count := 0
	res := make([]int, 2)
	for i, v := range nums {
		if count >= 2 {
			break
		}
		if v == head || v == tail {
			res[count] = i
			count++
		}
	}
	return res
}
