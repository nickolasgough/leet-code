package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead *ListNode
	var prevNode *ListNode
	currNode := head
	nextNode := currNode.Next
	nextNextNode := nextNode.Next
	for {
		nextNode.Next = currNode
		currNode.Next = nextNextNode
		if newHead == nil {
			newHead = nextNode
		}
		if prevNode != nil {
			prevNode.Next = nextNode
		}

		prevNode = currNode
		currNode = nextNextNode
		if nextNextNode == nil || nextNextNode.Next == nil {
			break
		}
		nextNode = nextNextNode.Next
		nextNextNode = nextNode.Next
	}
	return newHead
}
