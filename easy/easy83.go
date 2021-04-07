package easy

// 链表去重
func deleteDuplicate(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr, next := head, head.Next
	pre := &ListNode{Next: curr}
	res := pre

	for next != nil {
		if curr.Val == next.Val {
			pre.Next = next
			curr = next
			next = next.Next
			continue
		}
		pre, curr, next = pre.Next, curr.Next, next.Next
	}
	return res.Next
}
