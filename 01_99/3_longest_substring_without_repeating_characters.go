package _01_99

/*
Given a string s, find the length of the longest substring without repeating characters.


Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Example 4:
Input: s = ""
Output: 0


Constraints:
0 <= s.length <= 5 * 10^4
s consists of English letters, digits, symbols and spaces.

*/
func lengthOfLongestSubstring(s string) int {
	count := 0
	hash := make(map[byte]int)
	for i, j := 0, 0; j < len(s); j++ {
		i = max(i, hash[s[j]])
		count = max(count, j-i+1)
		hash[s[j]] = j + 1
	}
	return count
}
