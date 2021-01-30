package _300_399

/*
Given a positive integer n, break it into the sum of at least two positive integers and maximize
the product of those integers. Return the maximum product you can get.

Example 1:
Input: 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.

Example 2:
Input: 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.

Note: You may assume that n is not less than 2 and not larger than 58.

*/
func integerBreak(n int) int {
	if n < 2 {
		return 0
	}
	products := make([]int, n+1)
	products[0], products[1] = 1, 1
	for i := 2; i <= n; i++ {
		tmp := 0
		for j := 1; j < i; j++ {
			tmp = max(tmp, max(j*(i-j), j*products[i-j]))
		}
		products[i] = tmp
	}
	return products[n]
}
