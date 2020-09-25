package _100_199

import "container/list"

type Entry struct {
	k int
	v int
}

type LRUCache struct {
	hash  map[int]*list.Element
	list  *list.List
	total int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		hash:  make(map[int]*list.Element, capacity),
		list:  list.New(),
		total: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.hash[key]; ok {
		this.list.MoveToFront(e)
		return e.Value.(*Entry).v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.hash[key]; ok {
		e.Value.(*Entry).v = value
		this.list.MoveToFront(e)
		this.hash[key] = e
		return
	}
	if this.total == this.list.Len() {
		back := this.list.Back()
		this.list.Remove(back)
		delete(this.hash, back.Value.(*Entry).k)
	}
	e := &Entry{k: key, v: value}
	at := this.list.PushFront(e)
	this.hash[key] = at
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,val);
 */
