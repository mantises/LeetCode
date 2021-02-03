package _01_99

import "sort"

/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.



Example 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

Example 2:
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.


Constraints:
1 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^4

*/

func mergeInterval(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		merged := false
		for j := 0; j < len(res); j++ {
			if intervals[i][0] >= res[j][0] && intervals[i][0] <= res[j][1] {
				res[j][1] = max(intervals[i][1], res[j][1])
				merged = true
				break
			}
		}
		if !merged {
			res = append(res, intervals[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
