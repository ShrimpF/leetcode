package algo

import "fmt"

// ListNode ois
type ListNode struct {
	Val  int
	Next *ListNode
}

func printNode(head *ListNode) {
	for head.Next != nil {
		fmt.Print(head.Val, "=>")
		head = head.Next
	}
	fmt.Println(head.Val)
}

func reverseList1(head *ListNode) *ListNode {
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

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
