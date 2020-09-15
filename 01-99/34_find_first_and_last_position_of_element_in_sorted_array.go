package _01_99

func searchRangeByForce(nums []int, target int) []int {
	start, end := -1, -1
	for i, v := range nums {
		if v != target {
			continue
		}
		if start == -1 {
			start, end = i, i
		} else {
			end++
		}
	}
	return []int{start, end}
}

func searchRangeByBinarySearch(nums []int, target int) []int {
	start, end := -1, -1
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {

			for start = mid; start-1 >= 0 && nums[start-1] == target; {
				start--
			}
			for end = mid; end+1 <= len(nums)-1 && nums[end+1] == target; {
				end++
			}
			break
		}
	}
	return []int{start, end}
}
