package _400_499

import (
	"sort"
)

/*
Given a string, sort it in decreasing order based on the frequency of characters.

Example 1:
Input:
"tree"
Output:
"eert"
Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.

Example 2:
Input:
"cccaaa"
Output:
"cccaaa"
Explanation:
Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
Note that "cacaca" is incorrect, as the same characters must be together.

Example 3:
Input:
"Aabb"
Output:
"bbAa"
Explanation:
"bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.

*/

func frequencySort(s string) string {
	m := make(map[uint8]int)
	sli := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			sli = append(sli, s[i])
		}
		m[s[i]] += 1
	}
	sort.Slice(sli, func(i, j int) bool {
		return m[sli[i]] > m[sli[j]]
	})
	bytes := make([]byte, 0)
	for _, v := range sli {
		t := make([]byte, m[v])
		for i := 0; i < m[v]; i++ {
			t[i] = v
		}
		bytes = append(bytes, t...)
	}
	return string(bytes)
}
