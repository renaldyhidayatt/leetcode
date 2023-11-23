package on199

type LRUCache struct {
	head, tail *NodeLRU_Cache
	Keys       map[int]*NodeLRU_Cache
	Cap        int
}

type NodeLRU_Cache struct {
	Key, Val   int
	Prev, Next *NodeLRU_Cache
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Keys: make(map[int]*NodeLRU_Cache),
		Cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &NodeLRU_Cache{Key: key, Val: value}
		this.Keys[key] = node
		this.Add(node)
	}

	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *NodeLRU_Cache) {
	node.Prev = nil
	node.Next = this.head

	if this.head != nil {
		this.head.Prev = node
	}

	this.head = node

	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

func (this *LRUCache) Remove(node *NodeLRU_Cache) {
	if node == this.head {
		this.head = node.Next
		node.Next = nil
		return
	}

	if node == this.tail {
		this.tail = node.Prev
		node.Prev = nil
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
