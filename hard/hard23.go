package hard

func mergeKLists(lists []*ListNode) *ListNode {
	listss := splitList(lists)
	for _, v := range listss {
		ss := v
		go helper23(ss[0], ss[1])
	}
}

func splitList(lists []*ListNode) [][]*ListNode {
	res := make([][]*ListNode, 0, len(lists) / 2)
	for i := 0; i < len(lists); i += 2 {
		le, ri := i, i+2
		if len(lists) < ri {
			ri = len(lists)
		}
		res = append(res, lists[le:ri])
	}
	return res
}

// 合并两个有序链表
func helper23(l1 *ListNode, l2 *ListNode) *ListNode {
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