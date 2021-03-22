package medium

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, curr, next := &ListNode{}, head, head.Next
	pre.Next = curr
	res := pre
	var a, b, c, d *ListNode
	count := 1
	for curr != nil {
		if count == left {
			a, b = pre, curr
		}
		if count == right {
			c, d = curr, next
			break
		}
		pre, curr, next = pre.Next, curr.Next, next.Next
		count++
	}
	c.Next = nil
	reverse(b)
	a.Next = c
	b.Next = d
	return res.Next
}

func reverse(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var pre *ListNode
	curr := head
	next := head.Next
	for next != nil {
		curr.Next = pre
		pre = curr
		curr = next
		next = next.Next
	}
	curr.Next = pre
	return
}