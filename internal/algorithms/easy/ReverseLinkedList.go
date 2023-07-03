package easy

/*
Учитывая заголовок односвязного списка, переверните список и верните перевернутый список.
*/

/*
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {

	var cur, next, prev *ListNode

	cur = head

	for cur != nil {
		next = cur.Next // saving the next point
		cur.Next = prev // changing it previous
		prev = cur      // setting previous to current
		cur = next      // setting next to current
	}

	return prev
}

func ReverseList2(head *ListNode) (prev *ListNode) {

	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}

	return
}
