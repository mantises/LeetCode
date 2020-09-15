package _200_299

func moveZeroes(nums []int) {
	index := 0
	for _, v := range nums {
		if v != 0 {
			nums[index] = v
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}
