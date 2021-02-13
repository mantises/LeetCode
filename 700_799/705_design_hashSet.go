package _700_799

import "container/list"

/*
Design a HashSet without using any built-in hash table libraries.

Implement MyHashSet class:

void add(key) Inserts the value key into the HashSet.
bool contains(key) Returns whether the value key exists in the HashSet or not.
void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet,
do nothing.

Example 1:
Input
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]
Output
[null, null, null, true, false, null, true, null, false]

Explanation
MyHashSet myHashSet = new MyHashSet();
myHashSet.add(1);      // set = [1]
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(1); // return True
myHashSet.contains(3); // return False, (not found)
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(2); // return True
myHashSet.remove(2);   // set = [1]
myHashSet.contains(2); // return False, (already removed)


Constraints:
0 <= key <= 10^6
At most 10^4 calls will be made to add, remove, and contains.
Follow up: Could you solve the problem without using the built-in HashSet library?
*/

const SIZE = 1024

type MyHashSet struct {
	set []*list.List
}

/** Initialize your data structure here. */
func ConstructorMyHashSet() MyHashSet {
	s := make([]*list.List, SIZE)
	return MyHashSet{
		set: s,
	}
}

func (m *MyHashSet) Add(key int) {
	l := m.set[key%SIZE]
	if l == nil {
		l = list.New()
		m.set[key%SIZE] = l
	}
	for cnt, n := 0, l.Front(); cnt < l.Len(); {
		if n.Value.(int) == key {
			return
		}
		cnt++
		n = n.Next()
	}
	l.PushBack(key)
}

func (m *MyHashSet) Remove(key int) {
	l := m.set[key%SIZE]
	if l == nil {
		return
	}
	for n, cnt := l.Front(), 0; cnt < l.Len(); {
		if n.Value.(int) == key {
			l.Remove(n)
			return
		}
		cnt++
		n = n.Next()
	}
}

/** Returns true if this set contains the specified element */
func (m *MyHashSet) Contains(key int) bool {
	l := m.set[key%SIZE]
	if l == nil {
		return false
	}
	for cnt, n := 0, l.Front(); cnt < l.Len(); {
		if n.Value.(int) == key {
			return true
		}
		cnt++
		n = n.Next()
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
