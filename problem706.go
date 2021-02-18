package LeetCode

type MyHashMap struct {
	content int
	data     []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	capacity := 1000001
	data := make([]int, capacity)
	for i := range data {
		data[i] = -1
	}
	return MyHashMap{
		content: capacity,
		data:     data,
	}
}

func (this *MyHashMap) Hash(key int) int {
	return key % this.content
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.data[hm.Hash(key)] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.data[hm.Hash(key)]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.data[hm.Hash(key)] = -1
}