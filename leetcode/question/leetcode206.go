package question

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil {
		if slow == head {
			slow.Next = nil
		}
		tmp := fast.Next
		fast.Next = slow
		slow = fast
		fast = tmp
	}
	return slow
}
