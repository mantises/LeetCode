package _300_399

/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as
you like (ie, buy one and sell one share of the stock multiple times) with the following
restrictions:

You may not engage in multiple transactions at the same time (ie, you must sell the stock
before you buy again).
After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)

Example:
Input: [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]
*/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	holder := -prices[0] // holder with stock
	coolEmpty := 0       // empty and cool
	empty := 0           // empty and no cooldown
	for i := 1; i < len(prices); i++ {
		holderTmp := max(holder, empty-prices[i])
		coolEmptyTmp := holder + prices[i]
		emptyTmp := max(coolEmpty, empty)
		holder, coolEmpty, empty = holderTmp, coolEmptyTmp, emptyTmp
	}
	return max(empty, coolEmpty)
}
