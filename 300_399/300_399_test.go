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

func TestMaxCoins(t *testing.T) {
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums))
}

func TestMaxProfit(t *testing.T) {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}
