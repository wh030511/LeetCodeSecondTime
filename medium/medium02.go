package medium

// 链表合并
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = l2
	node := l2
	for ; l1 != nil && node != nil; l1, node = l1.Next, node.Next {
		node.Val += l1.Val
		pre = pre.Next
	}
	if node == nil {
		pre.Next = l1
	}

	f := 0
	pre = &ListNode{}
	pre.Next = l2
	for node := l2; node != nil; node = node.Next {
		num := node.Val + f
		node.Val = num % 10
		if num >= 10 {
			f = num / 10
		} else {
			f = 0
		}
		pre = pre.Next
	}

	if f != 0 {
		pre.Next = &ListNode{
			Val: f,
		}
	}

	return l2
}
