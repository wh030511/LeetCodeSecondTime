package medium

import (
	"fmt"
	"testing"
)

func TestMedium92(t *testing.T) {
	//param1 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val: 5,
	//				},
	//			},
	//		},
	//	},
	//}
	param2 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 5,
		},
	}

	res := reverseBetween(param2, 1, 2)
	fmt.Printf("res: %+#v", res)
}
