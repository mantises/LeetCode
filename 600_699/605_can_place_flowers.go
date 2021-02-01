package _600_699

/*
You have a long flowerbed in which some of the plots are planted, and some are not.
However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1
means not empty, and an integer n, return if n new flowers can be planted in the
flowerbed without violating the no-adjacent-flowers rule.


Example 1:
Input: flowerbed = [1,0,0,0,1], n = 1
Output: true

Example 2:
Input: flowerbed = [1,0,0,0,1], n = 2
Output: false


Constraints:
1 <= flowerbed.length <= 2 * 10^4
flowerbed[i] is 0 or 1.
There are no two adjacent flowers in flowerbed.
0 <= n <= flowerbed.length

*/

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; n > 0 && i < len(flowerbed); i += 2 {
		if flowerbed[i] == 0 {
			if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
				n--
			} else {
				i++
			}
		}
	}
	return n <= 0
}
