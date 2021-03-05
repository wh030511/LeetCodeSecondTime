package easy

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{}
	res := r
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			r.Next = l1
			l1 = l1.Next
		} else {
			r.Next = l2
			l2 = l2.Next
		}
		r = r.Next
	}
	if l1 != nil {
		r.Next = l1
	}
	if l2 != nil {
		r.Next = l2
	}
	return res.Next
}
