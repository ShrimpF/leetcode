package algo

import (
	"fmt"
)

// ListNode ois
type ListNode struct {
	Val  int
	Next *ListNode
}

// PrintNode is print the node's each Val
func PrintNode(head *ListNode) {
	for head.Next != nil {
		fmt.Print(head.Val, "=>")
		head = head.Next
	}
	fmt.Println(head.Val)
}

// ReverseList1 -- pick a node and connect
func ReverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := &ListNode{Val: head.Val}
	for head.Next != nil {
		node := &ListNode{Val: head.Next.Val}
		node.Next = result
		result = node
		head = head.Next
	}
	return result
}

// ReverseList2 -- recusive solution.
func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := ReverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
