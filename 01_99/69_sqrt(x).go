package _01_99

/*
Given a non-negative integer x, compute and return the square root of x.

Since the return type is an integer, the decimal digits are truncated, and only the
integer part of the result is returned.


Example 1:
Input: x = 4
Output: 2

Example 2:
Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since the decimal part is
truncated, 2 is returned.

Constraints:
0 <= x <= 2^31 - 1

*/
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left, right, mid := 1, x-1, 0
	for left <= right {
		mid = left + (right-left)>>1
		product := mid * mid
		if product > x {
			right = mid - 1
		} else if product < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right
}
