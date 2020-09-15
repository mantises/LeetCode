package _400_499

import "math"

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
