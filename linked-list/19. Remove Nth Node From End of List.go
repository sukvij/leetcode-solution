package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head, _ = solve(head, n)
	return head
}

func solve(head *ListNode, n int) (*ListNode, int) {
	if head == nil {
		return nil, 0
	}
	node, val := solve(head.Next, n)
	if val == n {
		return node, 1 + val
	}
	head.Next = node
	return head, 1 + val
}
