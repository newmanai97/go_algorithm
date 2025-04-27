package question

type MyLinkedList struct {
	Head *ListNode
	Size int
}

func Constructor() MyLinkedList {
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}

	return MyLinkedList{
		Head: head,
		Size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || index < 0 || index >= this.Size {
		return -1
	}

	cur := this.Head.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newhead := &ListNode{
		Val:  val,
		Next: this.Head.Next,
	}
	this.Head.Next = newhead
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	newtail := &ListNode{
		Val:  val,
		Next: nil,
	}
	cur.Next = newtail
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	} else if index > this.Size {
		return
	}

	newnode := &ListNode{
		Val: val,
	}
	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newnode.Next = cur.Next
	cur.Next = newnode
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	this.Size--
}
