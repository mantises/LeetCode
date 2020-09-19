package _200_299

import "testing"

func TestFindTheDuplicate(t *testing.T) {
	nums := []int{1, 3, 4, 2, 2}
	t.Log(findDuplicateByBinarySearch(nums))
}

func TestProduct(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	t.Log(productExceptSelf(nums))
}
