package _200_299

/*
Given a string s which represents an expression, evaluate this expression and return its value.

The integer division should truncate toward zero.

Example 1:
Input: s = "3+2*2"
Output: 7

Example 2:
Input: s = " 3/2 "
Output: 1
Example 3:

Input: s = " 3+5 / 2 "
Output: 5


Constraints:

1 <= s.length <= 3 * 10^5
s consists of integers and operators ('+', '-', '*', '/') separated by some number of spaces.
s represents a valid expression.
All the integers in the expression are non-negative integers in the range [0, 2^31 - 1].
The answer is guaranteed to fit in a 32-bit integer.

*/
func calculate(s string) int {
	nums := make([]int, 0)
	nums = append(nums, 0)
	perOperate := '+'
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			cur := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				cur = 10*cur + int(s[i]-'0')
			}
			i--
			switch perOperate {
			case '+':
				nums = append(nums, cur)
			case '-':
				nums = append(nums, -cur)
			case '*':
				pre := nums[len(nums)-1]
				nums = nums[:len(nums)-1]
				nums = append(nums, pre*cur)
			case '/':
				pre := nums[len(nums)-1]
				nums = nums[:len(nums)-1]
				nums = append(nums, int(pre/cur))
			}
		} else {
			perOperate = int32(s[i])
		}
	}
	for i := 1; i < len(nums); i++ {
		nums[0] += nums[i]
	}
	return nums[0]
}
