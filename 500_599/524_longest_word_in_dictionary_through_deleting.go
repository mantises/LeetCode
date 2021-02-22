package _500_599

import "sort"

/*
Given a string and a string dictionary, find the longest string in the dictionary
that can be formed by deleting some characters of the given string. If there are
more than one possible results, return the longest word with the smallest
lexicographical order. If there is no possible result, return the empty string.

Example 1:
Input:
s = "abpcplea", d = ["ale","apple","monkey","plea"]
Output:
"apple"

Example 2:
Input:
s = "abpcplea", d = ["a","b","c"]
Output:
"a"

Note:
All the strings in the input will only contain lower-case letters.
The size of the dictionary won't exceed 1,000.
The length of all the strings in the input won't exceed 1,000.

*/
func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		return len(d[i]) > len(d[j]) || (len(d[i]) == len(d[j]) && d[i] < d[j])
	})

	ret := ""
	for _, v := range d {
		if len(v) < len(ret) {
			continue
		}
		tmp := ""
		i, j := 0, 0
		for i < len(v) && j < len(s) {
			if v[i] == s[j] {
				j++
				i++
				if i == len(v) {
					tmp = v
				}
			} else {
				j++
			}
		}
		if tmp != "" && len(tmp) > len(ret) {
			ret = tmp
		}
	}
	return ret
}
