package _100_199

import "math"

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
