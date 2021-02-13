package _700_799

import (
	"container/list"
)

/*
Design a HashMap without using any built-in hash table libraries.

To be specific, your design should include these functions:

put(key, value) : Insert a (key, value) pair into the HashMap. If the value already exists in
the HashMap, update the value.
get(key): Returns the value to which the specified key is mapped, or -1 if this map contains
no mapping for the key.
remove(key) : Remove the mapping for the value key if this map contains the mapping for the key.

Example:
MyHashMap hashMap = new MyHashMap();
hashMap.put(1, 1);
hashMap.put(2, 2);
hashMap.get(1);            // returns 1
hashMap.get(3);            // returns -1 (not found)
hashMap.put(2, 1);          // update the existing value
hashMap.get(2);            // returns 1
hashMap.remove(2);          // remove the mapping for 2
hashMap.get(2);            // returns -1 (not found)

Note:
All keys and values will be in the range of [0, 1000000].
The number of operations will be in the range of [1, 10000].
Please do not use the built-in HashMap library.

*/

type Entry struct {
	k int
	v int
}

type MyHashMap struct {
	hash []*list.List
}

/** Initialize your data structure here. */
func ConstructorMyHashMap() MyHashMap {
	s := make([]*list.List, SIZE)
	return MyHashMap{
		hash: s,
	}
}

/** value will always be non-negative. */
func (m *MyHashMap) Put(key int, value int) {
	l := m.hash[key%SIZE]
	if l == nil {
		l = list.New()
		m.hash[key%SIZE] = l
	}
	for n, cnt := l.Front(), 0; cnt < l.Len(); {
		if n.Value.(*Entry).k == key {
			n.Value.(*Entry).v = value
			return
		}
		cnt++
		n = n.Next()
	}
	m.hash[key%SIZE].PushBack(&Entry{
		k: key,
		v: value,
	})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (m *MyHashMap) Get(key int) int {
	l := m.hash[key%SIZE]
	if l == nil {
		return -1
	}
	for n, cnt := l.Front(), 0; cnt < l.Len(); {
		if n.Value.(*Entry).k == key {
			return n.Value.(*Entry).v
		}
		cnt++
		n = n.Next()
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (m *MyHashMap) Remove(key int) {
	l := m.hash[key%SIZE]
	if l == nil {
		return
	}
	for n, cnt := l.Front(), 0; cnt < l.Len(); {
		if n.Value.(*Entry).k == key {
			l.Remove(n)
			return
		}
		cnt++
		n = n.Next()
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
