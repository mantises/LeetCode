package _01_99

import "math"

/*
Given two strings word1 and word2, return the minimum number of operations required to
convert word1 to word2.

You have the following three operations permitted on a word:
- Insert a character
- Delete a character
- Replace a character


Example 1:
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

Example 2:
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')


Constraints:
0 <= word1.length, word2.length <= 500
word1 and word2 consist of lowercase English letters.

*/
func minDistance(word1 string, word2 string) int {
	count := make([][]int, 0)
	for i := 0; i <= len(word1); i++ {
		count = append(count, make([]int, len(word2)+1))
	}
	for i := 1; i <= len(word2); i++ {
		count[0][i] = count[0][i-1] + 1
	}
	for i := 1; i <= len(word1); i++ {
		count[i][0] = count[i-1][0] + 1
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				count[i][j] = count[i-1][j-1]
			} else {
				count[i][j] = 1 + int(math.Min(float64(count[i-1][j-1]),
					math.Min(float64(count[i-1][j]), float64(count[i][j-1]))))
			}
		}
	}
	return count[len(word1)][len(word2)]
}

func minDistanceV2(word1 string, word2 string) int {
	count := make([]int, len(word2)+1)
	for i := 0; i <= len(word2); i++ {
		count[i] = i
	}
	var pre int
	for i := 1; i <= len(word1); i++ {
		temp := count[0]
		count[0] = i
		for j := 1; j <= len(word2); j++ {
			pre, temp = temp, count[j]
			if word1[i-1] == word2[j-1] {
				count[j] = pre
			} else {
				count[j] = 1 + int(math.Min(float64(count[j-1]),
					math.Min(float64(pre), float64(count[j]))))
			}
		}
	}
	return count[len(word2)]
}
