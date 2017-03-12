package linkedlist

import "errors"

// single linked list

type vtype interface{}

type llnode struct {
	val  vtype
	next *llnode
	pre  *llnode
}

type LinkedList struct {
	length uint32
	head   *llnode
	tail   *llnode
}

func AllocateLinkedList() *LinkedList {
	return &LinkedList{0, nil, nil}
}

func (l *LinkedList) Len() uint32 {
	return l.length
}

func (l *LinkedList) IsEmpty() bool {
	return l.length == 0
}

// add to the tail of the LL
func (l *LinkedList) Append(val vtype) {
	node := &llnode{val, nil, nil}
	if l.length == 0 {
		l.head = node
		l.tail = l.head
	} else {
		node.pre = l.tail
		l.tail.next = node
		l.tail = node

	}

	l.length++

}

// pop from the tail
func (l *LinkedList) Pop() (vtype, error) {
	if l.Len() == 0 {
		return nil, errors.New("Empty List!")
	}
	if l.Len() == 1 {
		ret := l.tail
		l.tail = nil
		l.head = nil
		l.length = 0
		return ret.val, nil
	}

	ret := l.tail
	l.tail = ret.pre
	l.tail.next = nil
	l.length--

	return ret.val, nil
}

// add to teh head of the LL
func (l *LinkedList) Push(val vtype) {
	node := &llnode{val, nil, nil}
	if l.length == 0 {
		l.head = node
		l.tail = l.head
	} else {
		node.next = l.head
		l.head.next = node
		l.head = node

	}

	l.length++

}

// pop from the head
func (l *LinkedList) Slice() (vtype, error) {
	if l.Len() == 0 {
		return nil, errors.New("Empty List!")
	}
	if l.Len() == 1 {
		ret := l.tail
		l.tail = nil
		l.head = nil
		l.length = 0
		return ret.val, nil
	}

	ret := l.head
	l.head = ret.next
	l.head.pre = nil
	l.length--

	return ret.val, nil
}

// compare_fn return true while a > b, else revert
func (l *LinkedList) SortLinkedList(ascending bool, compare_fn func(vtype, vtype) bool) {
	if l.Len() < 2 {
		return
	}

	// stupid bubblesort
	var swapped bool = true
	for swapped {

		curnode := l.head
		swapped = false

		for curnode.next != nil {
			var compare_result bool = compare_fn(curnode.val, curnode.next.val)

			if ascending == false {
				compare_result = !compare_result
			}

			if compare_result == true {
				tmp := curnode.val
				curnode.val = curnode.next.val
				curnode.next.val = tmp
				swapped = true
			}

			curnode = curnode.next

		}

	}

}
