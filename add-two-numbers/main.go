package main

// https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := l1
	node2 := l2
	carryOver := 0
	var headNode *ListNode
	var prevNode *ListNode
	for node1 != nil && node2 != nil {
		val := node1.Val + node2.Val + carryOver
		carryOver = 0

		if val >= 10 {
			carryOver = 1
			val -= 10
		}

		newNode := &ListNode{
			Val: val,
		}
		if headNode == nil {
			headNode = newNode
		} else if prevNode != nil {
			prevNode.Next = newNode
		}

		node1 = node1.Next
		node2 = node2.Next
		prevNode = newNode
	}

	if node1 != nil {
		prevNode.Next = node1
	} else if node2 != nil {
		prevNode.Next = node2
	}

	nextNode := prevNode.Next
	for carryOver > 0 {
		if nextNode == nil {
			nextNode = &ListNode{
				Val: carryOver,
			}
			prevNode.Next = nextNode
			break
		}

		nextNode.Val += carryOver
		carryOver = 0
		if nextNode.Val >= 10 {
			nextNode.Val -= 10
			carryOver = 1
		}
		prevNode = nextNode
		nextNode = nextNode.Next
	}

	return headNode
}
