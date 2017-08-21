package main

import (
	"fmt"
	"math"
)

func main() {
	l1 := makeList(5)
	l2 := makeList(5)
	res := addTwoNumbers2(l1, l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n := 0
	for i, e := 0, 2; ; i, e = i+1, 2 {
		if l1 != nil {
			n += int(math.Pow10(i)) * l1.Val
			l1 = l1.Next
		} else {
			e--
		}
		if l2 != nil {
			n += int(math.Pow10(i)) * l2.Val
			l2 = l2.Next
		} else {
			e--
		}
		if e == 0 {
			break
		}
	}
	head := &ListNode{}
	node := head
	for {
		node.Val = n % 10
		n /= 10
		if n > 0 {
			node.Next = &ListNode{}
			node = node.Next
		} else {
			break
		}
	}
	return head
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	node := head
	for {
		if l1 != nil {
			node.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			node.Val += l2.Val
			l2 = l2.Next
		}
		if node.Val > 9 {
			node.Next = &ListNode{}
			node.Next.Val = 1
			node.Val -= 10
		}
		if l1 != nil || l2 != nil {
			if node.Next == nil {
				node.Next = &ListNode{}
			}
		} else {
			break
		}
		node = node.Next
	}
	return head
}
func makeList(n int) *ListNode {
	head := &ListNode{}
	node := head
	for {
		node.Val = n % 10
		n /= 10
		if n > 0 {
			node.Next = &ListNode{}
			node = node.Next
		} else {
			break
		}
	}
	return head
}
