package main

import (
	"container/heap"
	"fmt"
	"time"
)

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}

type item struct {
	timestamp         int64
	key, value, index int
}

type priority_queue []*item

func (pq priority_queue) Len() int {
	return len(pq)
}

func (pq priority_queue) Less(i, j int) bool {
	return pq[i].timestamp < pq[j].timestamp
}

func (pq priority_queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priority_queue) Push(i any) {
	*pq = append((*pq), i.(*item))
}

func (pq *priority_queue) Pop() any {
	value := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return value
}

type LRUCache struct {
	cache map[int]*item
	queue priority_queue
	cap   int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache: make(map[int]*item),
		queue: make([]*item, 0),
		cap:   capacity}

	heap.Init(&lru.queue)
	return lru
}

func (l *LRUCache) Get(key int) int {
	mitem, ok := l.cache[key]
	if !ok {
		fmt.Printf("Item %d not found\n", key)
		return -1
	}

	mitem.timestamp = time.Now().UnixNano()
	heap.Fix(&l.queue, mitem.index)
	return mitem.value
}

func (l *LRUCache) Put(key, value int) {
	if item, ok := l.cache[key]; ok {
		item.value = value
		item.timestamp = time.Now().UnixNano()
		heap.Fix(&l.queue, item.index)
		return
	}

	if len(l.cache) == l.cap {
		toEvict := heap.Pop(&l.queue)
		delete(l.cache, toEvict.(*item).key)
		fmt.Printf("Item %d evicted from the cache\n", toEvict.(*item).key)
	}

	item := item{time.Now().UnixNano(), key, value, len(l.cache)}
	heap.Push(&l.queue, &item)
	l.cache[key] = &item
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
