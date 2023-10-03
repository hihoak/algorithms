package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}

	mergedList := swapPairs(list1)

	for mergedList != nil {
		fmt.Printf("%d ", mergedList.Val)
		mergedList = mergedList.Next
	}
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	result := head
	if head.Next != nil {
		result = head.Next
	}

	var tempHead *ListNode
	var previousHead *ListNode

	for head.Next != nil && head.Next.Next != nil {
		tempHead = head
		head = head.Next
		if previousHead != nil {
			previousHead.Next = head
		}
		tempHead.Next = head.Next
		head.Next = tempHead

		previousHead = head.Next
		head = head.Next.Next
	}

	return result
}
