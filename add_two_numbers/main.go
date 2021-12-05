// https://leetcode.com/problems/add-two-numbers/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}}
	res := addTwoNumbers(l1, l2)

	fmt.Printf("Result is %s", res.Print())
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) Print() string {
	resSlice := make([]string, 0)
	for l != nil {
		resSlice = append(resSlice, strconv.Itoa(l.Val))
		l = l.Next
	}
	return strings.Join(resSlice, " - ")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resList := &ListNode{}
	curList := resList
	remainder := 0
	for l1 != nil || l2 != nil {
		val := remainder
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		curList.Next = &ListNode{Val: val % 10}
		remainder = val / 10
		curList = curList.Next
	}

	if remainder != 0 {
		curList.Next = &ListNode{Val: remainder}
	}
	return resList.Next
}
