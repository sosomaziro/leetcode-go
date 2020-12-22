package main

func main() {
	var l1 ListNode
	var l11 ListNode
	var l12 ListNode
	l1.Val = 1
	l11.Val = 8
	l12.Val = 5
	l1.Next = &l11
	l11.Next = &l12
	var l2 ListNode
	var l21 ListNode
	var l22 ListNode
	var l23 ListNode
	var l24 ListNode
	l2.Val = 9
	l21.Val = 8
	l22.Val = 4
	l23.Val = 9
	l24.Val = 9
	l2.Next = &l21
	l21.Next = &l22
	l22.Next = &l23
	l23.Next = &l24
	numbers := addTwoNumbers(&l1, &l2)
	print(numbers)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := getSum(l1, l2, 0)
	return root
}

func getSum(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	var result ListNode
	if l1 == nil && l2 == nil {
		if carry == 0 {
			return nil
		} else {
			result.Val = carry
			return &result
		}
	}
	var v1 int
	var v2 int
	var s1 *ListNode
	var s2 *ListNode
	if l1 == nil {
		v1 = 0
		s1 = nil
	} else {
		v1 = l1.Val
		s1 = l1.Next
	}
	if l2 == nil {
		v2 = 0
		s2 = nil
	} else {
		v2 = l2.Val
		s2 = l2.Next
	}
	val := v1 + v2 + carry
	if val >= 10 {
		val = val - 10
		carry = 1
	} else {
		carry = 0
	}
	result.Next = getSum(s1, s2, carry)
	result.Val = val
	return &result
}

func print(list *ListNode)  {
	if list != nil {
		println(list.Val)
		print(list.Next)
	}
}