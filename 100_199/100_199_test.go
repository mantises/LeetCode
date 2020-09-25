package _100_199

import "testing"

func TestMajorityElementMooreVote(t *testing.T) {
	nums := []int{2, 3, 4, 3, 2, 2, 5, 2, 2}
	t.Log(majorityElementMooreVote(nums))
}

func TestLRUCache(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(2, 6)
	lRUCache.Put(1, 5)
	t.Log(lRUCache.Get(1)) // return -1 (not found)

	lRUCache.Put(1, 2) // evicts key 1

	t.Log(lRUCache.Get(1)) // return -1 (not found)
	t.Log(lRUCache.Get(2)) // return 3
}
