package _100_199

import "math"

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in
constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.

Example 1:
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
*/

type MinStack struct {
	nums []int
	min  []int
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{nums: make([]int, 0), min: make([]int, 0)}
}

func (m *MinStack) Push(x int) {
	if len(m.min) == 0 || (len(m.min) != 0 && x <= m.min[len(m.min)-1]) {
		m.min = append(m.min, x)
	}
	m.nums = append(m.nums, x)
}

func (m *MinStack) Pop() {
	if len(m.nums) == 0 {
		return
	}
	if m.nums[len(m.nums)-1] == m.min[len(m.min)-1] {
		m.min = m.min[:len(m.min)-1]
	}
	m.nums = m.nums[:len(m.nums)-1]
}

func (m *MinStack) Top() int {
	if len(m.nums) == 0 {
		return math.MinInt32
	}
	return m.nums[len(m.nums)-1]
}

func (m *MinStack) GetMin() int {
	if len(m.min) == 0 {
		return math.MinInt32
	}
	return m.min[len(m.min)-1]
}
