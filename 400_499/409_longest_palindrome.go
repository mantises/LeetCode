package _400_499

/*
Given a string s which consists of lowercase or uppercase letters, return the length
of the longest palindrome that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome here.


Example 1:
Input: s = "abccccdd"
Output: 7
Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.

Example 2:
Input: s = "a"
Output: 1

Example 3:
Input: s = "bb"
Output: 2


Constraints:
1 <= s.length <= 2000
s consists of lowercase and/or uppercase English letters only.

*/
func longestPalindrome(s string) int {
	hash := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}
	maxLength, tmp := 0, 0
	for _, v := range hash {
		if v%2 == 0 {
			maxLength += v
		} else {
			tmp = max(tmp, v)
			maxLength += v - 1
		}
	}
	if tmp > 0 {
		maxLength += 1
	}
	return maxLength
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
