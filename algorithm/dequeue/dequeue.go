package dequeue

import "container/list"

type Deque struct {
	items *list.List
}

func NewDeque() *Deque {
	return &Deque{items: list.New()}
}

func (d *Deque) IsEmpty() bool {
	return d.items.Len() == 0
}

func (d *Deque) AddRear(item interface{}) {
	d.items.PushBack(item)
}

func (d *Deque) AddFront(item interface{}) {
	d.items.PushFront(item)
}

func (d *Deque) RemoveFront() interface {
} {
	front := d.items.Front()

	if front != nil {
		d.items.Remove(front)

		return front.Value
	}

	return nil
}

func (d *Deque) RemoveRear() interface{} {
	back := d.items.Back()

	if back != nil {
		d.items.Remove(back)
		return back.Value
	}

	return nil

}

func (d *Deque) Size() int {
	return d.items.Len()
}

func DequeueMain() {
	d := NewDeque()
	println(d.IsEmpty())
	d.AddRear(8)
	d.AddRear(5)
	d.AddFront(7)
	d.AddFront(10)
	println(d.Size())
	println(d.IsEmpty())
	d.AddRear(11)
	println(d.RemoveRear())
	println(d.RemoveFront())
	d.AddFront(55)
	d.AddRear(45)
	for e := d.items.Front(); e != nil; e = e.Next() {
		println(e.Value)
	}
}
