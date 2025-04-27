package question

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temphead := &ListNode{
		Next: head,
	}
	father := temphead
	node1 := head
	node2 := head.Next
	for node1 != nil && node2 != nil {
		father.Next = node2
		temp := node2.Next
		node2.Next = node1
		node1.Next = temp
		father = node1
		node1 = temp
		if temp != nil {
			node2 = temp.Next
		} else {
			node2 = nil
		}
	}
	return temphead.Next
}
