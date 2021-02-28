package _01_99

/*
Given the head of a singly linked list and two integers left and right where left <=
right, reverse the nodes of the list from position left to position right, and return
the reversed list.


Example 1:
Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]

Example 2:
Input: head = [5], left = 1, right = 1
Output: [5]


Constraints:
The number of nodes in the list is n.
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	newHead := &ListNode{Next: head}
	lHead, rHead := newHead, newHead
	cnt := 1
	for {
		if cnt < left {
			lHead = lHead.Next
		}
		if cnt <= right {
			cnt++
			rHead = rHead.Next
		} else {
			break
		}
	}
	subhead := lHead.Next
	lastTail := rHead.Next
	rHead.Next = nil
	subhead, subtail := reverse(subhead)
	lHead.Next = subhead
	subtail.Next = lastTail
	return newHead.Next
}

func reverse(l *ListNode) (*ListNode, *ListNode) {
	if l == nil || l.Next == nil {
		return l, l
	}
	tail := l
	p, q := l, l.Next
	l.Next = nil
	for q != nil {
		t := q.Next
		q.Next = p
		p = q
		q = t
	}
	return p, tail
}
