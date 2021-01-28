package reverselinkedlist

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func List2Integers(head *ListNode) []int {
	// panic when list length beyond limit
	limit := 100

	times := 0

	// var res = []int represents nil
	// res := []int{} represents []
	// res := make([]int, 0) represents []
	res := make([]int, 0)

	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("Beyond list print length %d, please check or adjust the limit in func List2Integers", limit)
			panic(msg)
		}

		res = append(res, head.value)
		head = head.next
	}

	return res
}

func InitLinkedList(initials []int) *ListNode {
	if len(initials) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l // t is the tmp variable to represent the current node
	for _, v := range initials {
		t.next = &ListNode{value: v}
		t = t.next
	}

	return l.next
}

func (head *ListNode) Reverse() *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.next
		head.next = prev
		prev = head
		head = next
	}

	fmt.Println("Reversed linked list: ", List2Integers(prev))

	return prev
}
