package go_code

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type linkData struct {
	key, val int
	prev *linkData
	next *linkData
}


type LRUCache struct {
	head *linkData
	tail *linkData

	cap int
	hashmap map[int]*linkData
}


func Constructor(capacity int) LRUCache {
	head := &linkData{0,0, nil,nil}
	tail := &linkData{0,0, nil,nil}
	head.next = tail
	tail.prev = head
	return LRUCache{head, tail,  capacity, make(map[int]*linkData)}
}


func (this *LRUCache) Get(key int) int {
	item, ok := this.hashmap[key]
	if !ok{
		return -1
	}

	// 最新元素move to head
	this.moveToHead(item)
	return item.val
}


func (this *LRUCache) Put(key int, value int)  {
	last := this.tail.prev
	item, ok := this.hashmap[key]
	if ok{
		item.val = value
		this.moveToHead(item)
	}else{
		// 淘汰末尾
		if len(this.hashmap) >= this.cap{
			this.remove(last)
			delete(this.hashmap, last.key)
		}
		node := &linkData{key, value, nil, nil}
		this.hashmap[key] = node
		this.addToHead(node)
	}
}

// 新增元素放头部
func (this *LRUCache)addToHead(data *linkData)  {
	data.next = this.head.next
	this.head.next.prev = data
	this.head.next = data
	data.prev =this.head
}

func (this *LRUCache)moveToHead(data *linkData)  {
	this.remove(data)
	this.addToHead(data)
}

func (this *LRUCache)remove(data *linkData)  {
	data.prev.next = data.next
	data.next.prev = data.prev
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/*
LRUCache cache = new LRUCache( 2 );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得关键字 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得关键字 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4
 */
func TestLRUCache(t *testing.T)  {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2,2)
	assert.Equal(t, lru.Get(1), 1)
	lru.Put(3,3)
	assert.Equal(t, -1, lru.Get(2), )
	lru.Put(4,4)
	assert.Equal(t, -1, lru.Get(1), )
	assert.Equal(t, 3, lru.Get(3), )
	assert.Equal(t, 4, lru.Get(4), )



}
