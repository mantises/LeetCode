package _200_299

func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	for i := 0; i < len(nums); i++ {
		res[i] = 1
	}
	left, right := 1, 1
	for i := 0; i < length; i++ {
		res[i] *= left
		left *= nums[i]

		res[length-1-i] *= right
		right *= nums[length-1-i]
	}
	return res
}
