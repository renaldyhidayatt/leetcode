package on299

type MyStack struct {
	enque []int
	deque []int
}

func Constructor225() MyStack {
	return MyStack{[]int{}, []int{}}
}

func (this *MyStack) Push(x int) {
	this.enque = append(this.enque, x)
}

func (this *MyStack) Pop() int {
	length := len(this.enque)

	for i := 0; i < length-1; i++ {
		this.deque = append(this.deque, this.enque[0])
		this.enque = this.enque[1:]
	}

	topEle := this.enque[0]
	this.enque = this.deque
	this.deque = nil

	return topEle
}

func (this *MyStack) Top() int {
	topEl := this.Pop()

	this.enque = append(this.enque, topEl)

	return topEl
}

func (this *MyStack) Empty() bool {
	if len(this.enque) == 0 {
		return true
	}

	return false
}
