package _300_399

import (
	"math"
	"sort"
)

/*
You are given coins of different denominations and a total amount of money amount. Write a
function to compute the fewest number of coins that you need to make up that amount. If that
amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

Example 1:
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:
Input: coins = [2], amount = 3
Output: -1

Example 3:
Input: coins = [1], amount = 0
Output: 0

Example 4:
Input: coins = [1], amount = 1
Output: 1

Example 5:
Input: coins = [1], amount = 2
Output: 2

Constraints:
1 <= coins.length <= 12
1 <= coins[i] <= 2^31 - 1
0 <= amount <= 10^4

*/

func coinChange(coins []int, amount int) int {
	count := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		count[i] = amount + 1
	}
	sort.Ints(coins)
	count[0] = 0
	for _, v := range coins {
		for i := v; i <= amount; i++ {
			count[i] = int(math.Min(float64(count[i]), float64(1+count[i-v])))
		}
	}
	if count[amount] == amount+1 {
		return -1
	}
	return count[amount]
}
