package question

type Stack []int

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}

func (s *Stack) Pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

type MyQueue struct {
	stackin  *Stack
	stackout *Stack
}

//func Constructor() MyQueue {
//	return MyQueue{
//		stackin:  &Stack{},
//		stackout: &Stack{},
//	}
//}

func (this *MyQueue) Push(x int) {
	this.stackin.Push(x)
}

func (this *MyQueue) Pop() int {
	this.fillStackOut()
	return this.stackout.Pop()
}

func (this *MyQueue) Peek() int {
	this.fillStackOut()
	return this.stackout.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.stackin.Empty() && this.stackout.Empty()
}

func (this *MyQueue) fillStackOut() {
	if this.stackout.Empty() {
		for !this.stackin.Empty() {
			val := this.stackin.Pop()
			this.stackout.Push(val)
		}
	}
}
