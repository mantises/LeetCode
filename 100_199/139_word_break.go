package _100_199

/*
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
determine if s can be segmented into a space-separated sequence of one or more dictionary
words.

Note:
The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.

Example 1:
Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.

Example 3:
Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false

*/
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	res := make([]bool, len(s)+1)
	res[0] = true
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			if _, ok := wordMap[s[i:j]]; ok && res[i] {
				res[j] = true
			}
		}
	}
	return res[len(s)]
}
