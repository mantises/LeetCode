package _200_299

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false


Constraints:
1 <= s.length, t.length <= 5 * 10^4
s and t consist of lowercase English letters.


Follow up: What if the inputs contain Unicode characters?
	How would you adapt your solution to such a case?

*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a, b := [26]int{}, [26]int{}
	for i := 0; i < len(s); i++ {
		a[s[i]-'a'] += 1
		b[t[i]-'a'] += 1
	}
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
