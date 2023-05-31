package main

import "fmt"

// ListNode is a node of a linked list.
type ListNode struct {
	Next *ListNode
	Val  int
}


func main() {
	list := &ListNode{
		Next: &ListNode{
			Next: nil,
			Val:  20,
		},
	}

	fmt.Println(list)
}

// BEGIN (write your solution here)
func (head *ListNode) Reverse() *ListNode {

	if head == nil {
		return nil
	}

	var r *ListNode

	curr := head
	for curr != nil {
		r = &ListNode{
			Next: r,
			Val:  curr.Val,
		}
		curr = curr.Next
	}

	return r

}
