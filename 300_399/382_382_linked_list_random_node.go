package _300_399

import "math/rand"

/*
Given a singly linked list, return a random node's value from the linked list. Each
node must have the same probability of being chosen.


Example 1:
Input
["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
[[[1, 2, 3]], [], [], [], [], []]
Output
[null, 1, 3, 2, 2, 3]

Explanation
Solution solution = new Solution([1, 2, 3]);
solution.getRandom(); // return 1
solution.getRandom(); // return 3
solution.getRandom(); // return 2
solution.getRandom(); // return 2
solution.getRandom(); // return 3
// getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.


Constraints:
The number of nodes in the linked list will be in the range [1, 104]
-10^4 <= Node.val <= 10^4
At most 104 calls will be made to getRandom.

Follow up:
What if the linked list is extremely large and its length is unknown to you?
Could you solve this efficiently without using extra space?

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	l      *ListNode
	length int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	s := Solution{l: &ListNode{Next: head}}
	for t := s.l; t != nil; s.length++ {
		t = t.Next
	}
	return s
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	r := rand.Intn(this.length)
	t := this.l
	for r != 0 {
		t = t.Next
	}
	return t.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
