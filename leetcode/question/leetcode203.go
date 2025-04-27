package question

func removeElements(head *ListNode, val int) *ListNode {
	temphead := &ListNode{}
	temphead.Next = head

	cur := temphead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return temphead.Next

}

func removeElements1(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return cur
}
