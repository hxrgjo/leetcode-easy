package main

func deleteDuplicates(head *ListNode) *ListNode {

	result := new(ListNode)
	temp := result
	m := make(map[int]interface{})

	for head != nil {

		if _, ok := m[head.Val]; !ok {
			var i interface{}
			m[head.Val] = i
			temp.Next = new(ListNode)
			temp.Next.Val = head.Val
			temp = temp.Next
		}

		head = head.Next
	}

	return result.Next
}
