package _400_499

import (
	"math"
)

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		index := int(math.Abs(float64(v))) - 1
		if nums[index] > 0 && index < len(nums) {
			nums[index] = -nums[index]
		}
	}
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			nums[index] = i + 1
			index++
		}
	}
	return nums[:index]
}

func findDisappearedNumbersII(nums []int) []int {
	n := len(nums)
	for _, v := range nums {
		index := (v - 1) % n
		nums[index] += n
	}
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if nums[i] <= n {
			res = append(res, i+1)
		}
	}
	return res
}
