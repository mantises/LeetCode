package _300_399

/*
Given a string s and an integer k, return the length of the longest substring of s
 that contains at most k distinct characters.

Example 1:
Input: s = "eceba", k = 2
Output: 3
Explanation: The substring is "ece" with length 3.

Example 2:
Input: s = "aa", k = 1
Output: 2
Explanation: The substring is "aa" with length 2.

Constraints:
1 <= s.length <= 5 * 10^4
0 <= k <= 50
*/
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	ret := 0
	hash := make(map[uint8]int)
	for i, j := 0, 0; j < len(s); j++ {
		hash[s[j]] += 1
		if len(hash) <= k {
			ret = max(ret, j-i+1)
		} else {
			for len(hash) > k {
				hash[s[i]] -= 1
				if hash[s[i]] == 0 {
					delete(hash, s[i])
				}
				i++
			}
		}
	}
	return ret
}
