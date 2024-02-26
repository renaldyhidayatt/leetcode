package cache

import "container/list"

type item struct {
	key   string
	value any
	freq  int
}

type LFU struct {
	len     int
	cap     int
	minFreq int
	itemMap map[string]*list.Element
	freqMap map[int]*list.List
}

func NewLFU(capacity int) *LFU {
	return &LFU{
		len:     0,
		cap:     capacity,
		minFreq: 0,
		itemMap: make(map[string]*list.Element),
		freqMap: make(map[int]*list.List),
	}
}

func initItem(k string, v any, f int) item {
	return item{
		key:   k,
		value: v,
		freq:  f,
	}
}

func (c *LFU) increaseFreq(e *list.Element) {
	obj := e.Value.(item)

	oldLost := c.freqMap[obj.freq]

	oldLost.Remove(e)

	if c.minFreq == obj.freq && oldLost.Len() == 0 {
		c.minFreq++
	}

	c.insertMap(obj)
}

func (c *LFU) Get(key string) any {
	if e, ok := c.itemMap[key]; ok {
		c.increaseFreq(e)
		obj := e.Value.(item)

		return obj.value
	}

	return nil
}

func (c *LFU) Put(key string, value any) {
	if e, ok := c.itemMap[key]; ok {
		obj := e.Value.(item)
		obj.value = value
		c.increaseFreq(e)
	} else {
		obj := initItem(key, value, 1)

		if c.len == c.cap {
			c.eliminate()
			c.len--
		}

		c.insertMap(obj)

		c.minFreq = 1
		c.len++
	}
}

func (c *LFU) insertMap(obj item) {
	l, ok := c.freqMap[obj.freq]

	if !ok {
		l = list.New()
		c.freqMap[obj.freq] = l
	}

	e := l.PushFront(obj)

	c.itemMap[obj.key] = e
}

func (c *LFU) eliminate() {
	l := c.freqMap[c.minFreq]
	e := l.Back()
	obj := e.Value.(item)

	l.Remove(e)

	delete(c.itemMap, obj.key)
}
