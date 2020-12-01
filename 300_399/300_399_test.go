package _300_399

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{2, 3, 5}
	fmt.Println(coinChange(coins, 15))
}

func TestLengthOfLIS(t *testing.T) {
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(nums))
}
