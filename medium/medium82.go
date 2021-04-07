package medium

// 链表去重2.0
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	record := make(map[int]bool)
	curr, next := head, head.Next

	for next != nil { // 记录所有重复的数字
		if curr.Val == next.Val {
			record[curr.Val] = true
		}
		curr, next = curr.Next, next.Next
	}

	curr, next = head, head.Next
	pre := &ListNode{Next: curr}
	res := pre
	for next != nil {
		if _, ok := record[curr.Val]; ok {
			curr, next = next, next.Next
			pre.Next = curr
			continue
		}
		pre, curr, next = pre.Next, curr.Next, next.Next
	}
	if _, ok := record[curr.Val]; ok {
		pre.Next = next
	}
	return res.Next
}
