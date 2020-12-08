package _400_499

import "testing"

func TestFindDisappearedNumbers(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	t.Log(findDisappearedNumbers(nums))
}

func TestPartitionEqualSubsetSum(t *testing.T) {
	nums := []int{5, 3, 9, 7, 4}
	t.Log(canPartition(nums))
}
