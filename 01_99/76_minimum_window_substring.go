package _01_99

/*
Given two strings s and t, return the minimum window in s which will contain all the
characters in t. If there is no such window in s that covers all characters in t, return
the empty string "".

Note that If there is such a window, it is guaranteed that there will always be only one
unique minimum window in s.


Example 1:
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"

Example 2:
Input: s = "a", t = "a"
Output: "a"


Constraints:
1 <= s.length, t.length <= 105
s and t consist of English letters.

Follow up: Could you find an algorithm that runs in O(n) time?
*/
func minWindow(s string, t string) string {
	pattenMap := make(map[uint8]int)
	for _, v := range t {
		pattenMap[uint8(v-'A')] += 1
	}
	cnt := 0
	currentMap := make(map[uint8]int)
	ret, tmp := "", ""
	for i, j := 0, 0; j < len(s); j++ {
		currentMap[s[j]-'A'] += 1
		if pattenMap[s[j]-'A'] > 0 && currentMap[s[j]-'A'] <= pattenMap[s[j]-'A'] {
			cnt++
		}
		if cnt == len(t) {
			tmp = s[i : j+1]
			for i < j && currentMap[s[i]-'A'] > pattenMap[s[i]-'A'] {
				currentMap[s[i]-'A'] -= 1
				i++
				tmp = s[i : j+1]
			}
			if ret == "" || len(tmp) < len(ret) {
				ret = tmp
			}
		}
	}
	return ret
}
