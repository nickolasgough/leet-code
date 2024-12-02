package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	tail := head
	n := 1
	for n < k && tail != nil {
		tail = tail.Next
		n += 1
	}
	if n < k || tail == nil {
		return head
	}

	nextHead := reverseKGroup(tail.Next, k)
	tail.Next = nil

	newHead, newTail := reverseList(head)
	newTail.Next = nextHead
	return newHead
}

func reverseList(head *ListNode) (*ListNode, *ListNode) {
	var prevNode *ListNode
	currNode := head
	tail := head
	for currNode != nil {
		nextNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}
	return prevNode, tail
}
