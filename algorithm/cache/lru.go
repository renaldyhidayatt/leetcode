package cache

import "algoritmAndDs/structure/linkedlist"

type item_lru struct {
	key   string
	value any
	freq  int
}

type LRU struct {
	dl       *linkedlist.Doubly[any]
	size     int
	capacity int
	storage  map[string]*linkedlist.Node[any]
}

func NewLRU(capacity int) LRU {
	return LRU{
		dl:       linkedlist.NewDoubly[any](),
		storage:  make(map[string]*linkedlist.Node[any], capacity),
		size:     0,
		capacity: capacity,
	}
}

func (c *LRU) Get(key string) any {

	if e, ok := c.storage[key]; ok {
		c.dl.MoveToBack(e)
		return e.Val.(item).value
	}
	return nil
}

func (c *LRU) Put(key string, value any) {
	e, ok := c.storage[key]
	if ok {
		n := e.Val.(item)
		n.value = value
		e.Val = n
		c.dl.MoveToBack(e)
		return
	}

	if c.size >= c.capacity {
		e := c.dl.Front()
		dk := e.Val.(item).key
		c.dl.Remove(e)
		delete(c.storage, dk)
		c.size--
	}

	n := item_lru{key: key, value: value}
	c.dl.AddAtEnd(n)
	ne := c.dl.Back()
	c.storage[key] = ne
	c.size++
}
