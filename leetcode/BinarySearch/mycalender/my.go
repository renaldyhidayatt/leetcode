package mycalender

type Event struct {
	start, end  int
	left, right *Event
}

func (e *Event) Insert(curr *Event) bool {
	if e.end > curr.start && curr.end > e.start {
		return false
	}

	if curr.start < e.start {
		if e.left == nil {
			e.left = curr
		} else {
			return e.left.Insert(curr)
		}
	} else {
		if e.right == nil {
			e.right = curr
		} else {
			return e.right.Insert(curr)
		}
	}

	return true
}

type MyCalendar struct {
	root *Event
}

func Constructor() MyCalendar {
	return MyCalendar{
		root: nil,
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	curr := &Event{start: start, end: end, left: nil, right: nil}
	if this.root == nil {
		this.root = curr

		return true
	}

	return this.root.Insert(curr)
}
