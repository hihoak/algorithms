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
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	list2 := &ListNode{Val: -10, Next: &ListNode{Val: 0, Next: &ListNode{Val: 30, Next: nil}}}

	mergedList := mergeTwoLists(list1, list2)

	for mergedList != nil {
		fmt.Printf("%d ", mergedList.Val)
		mergedList = mergedList.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return result.Next
}
