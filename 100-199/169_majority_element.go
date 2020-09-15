package _100_199

import "sort"

func majorityElementBySort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)>>1]
}

func majorityElementByHashMap(nums []int) int {
	m := make(map[int]int)
	max := 0
	var val int
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
		if m[v] >= max {
			max = m[v]
			val = v
		}
	}
	return val
}

func majorityElementMooreVote(nums []int) int {
	var candidate, count int
	for _, v := range nums {
		if count == 0 {
			candidate = v
			count++
			continue
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
