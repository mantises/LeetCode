package _400_499

import (
	"sort"
)

/*
Given a collection of intervals, find the minimum number of intervals you need to remove
to make the rest of the intervals non-overlapping.

Example 1:
Input: [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.

Example 2:
Input: [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.

Example 3:
Input: [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.

Note:
You may assume the interval's end point is always bigger than its start point.
Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.

*/
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	lastInterval := intervals[0]
	count := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= lastInterval[1] {
			count++
			lastInterval = intervals[i]
		}
	}
	return len(intervals) - count
}
