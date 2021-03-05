package medium

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil{
		return nil
	}
	if head.Next.Next == nil {
		if n == 1 {
			head.Next = nil
			return head
		}
		return head.Next
	}
	f := head
	for ; n > 0; n-- {
		f = f.Next
	}
	if f == nil {
		return head.Next
	}
	m := head.Next.Next
	l := head
	for f.Next != nil {
		f = f.Next
		m = m.Next
		l = l.Next
	}
	l.Next = m
	return head
}
