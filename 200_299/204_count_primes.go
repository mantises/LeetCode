package _200_299

/*
Count the number of prime numbers less than a non-negative number, n.

Example 1:
Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.

Example 2:
Input: n = 0
Output: 0

Example 3:
Input: n = 1
Output: 0

Constraints:
0 <= n <= 5 * 10^6

*/
func countPrimes(n int) int {
	counter := make([]int8, n+1)
	counter[0], counter[1] = 1, 1
	for i := 2; i <= n/2; i++ {
		for j := 2; j*i <= n; j += 1 {
			counter[i*j] = 1
		}
	}
	res := 0
	for i := 0; i <= n; i++ {
		if counter[i] == 0 {
			res++
		}
	}
	return res
}
