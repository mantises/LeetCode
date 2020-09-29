package _01_99

func sortColors(nums []int) {
	low, cur, high := 0, 0, len(nums)-1
	for cur <= high {
		if nums[cur] == 0 {
			nums[cur], nums[low] = nums[low], 0
			low++
			cur++
		} else if nums[cur] == 2 {
			nums[cur], nums[high] = nums[high], 2
			high--
		} else {
			cur++
		}
	}
}
