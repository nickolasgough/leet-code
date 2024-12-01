package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = reverseList(head)
	head = removeNthFromStart(head, n)
	return reverseList(head)
}

func removeNthFromStart(head *ListNode, n int) *ListNode {
	i := 1
	currNode := head
	var prevNode *ListNode
	for i != n && currNode != nil {
		i += 1
		prevNode = currNode
		currNode = currNode.Next
	}
	if currNode == nil {
		return head
	}
	if prevNode == nil {
		return currNode.Next
	}
	prevNode.Next = currNode.Next
	return head
}

func reverseList(head *ListNode) *ListNode {
	currNode := head
	var prevNode *ListNode
	for currNode != nil {
		nextNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}
	return prevNode
}
