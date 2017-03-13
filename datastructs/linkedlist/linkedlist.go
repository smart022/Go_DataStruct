package linkedlist

import "errors"

// double oriented linked list

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

//  LinkedList Iterator
type ll_iter struct {
	list *LinkedList
	node *llnode
}

func (l *LinkedList) LLMakeIterator() (*ll_iter, error) {
	if l.IsEmpty() {
		return nil, errors.New("Empty LinkedList!")
	}

	iter := new(ll_iter)
	iter.list = l
	iter.node = l.head

	return iter, nil
}

func (iter *ll_iter) LLIteratorHasNext() bool {
	if iter.node.next != nil {
		return true
	} else {
		return false
	}
}

func (iter *ll_iter) LLIteratorHasPrev() bool {
	if iter.node.pre != nil {
		return true
	} else {
		return false
	}
}

func (iter *ll_iter) LLIteratorNext() error {
	if iter.LLIteratorHasNext() {
		iter.node = iter.node.next
		return nil
	} else {
		return errors.New("already met the tail,No elements left!")
	}

}

func (iter *ll_iter) LLIteratorPrev() error {
	if iter.LLIteratorHasPrev() {
		iter.node = iter.node.pre
		return nil
	} else {
		return errors.New("already at the begining,No elements forward!")
	}

}

func (l *ll_iter) LLIteratorGetPayload() (vtype, error) {
	if l == nil || l.node == nil {
		return nil, errors.New("Invalid Iterator!")
	}

	return l.node.val, nil
}

// Delete the node the iterator is pointing to.  After deletion, the iterator:
//
// - is invalid and cannot be used (but must be freed), if there was only one
//   element in the list
//
// - the successor of the deleted node, if there is one.
//
// - the predecessor of the deleted node, if the iterator was pointing at
//   the tail.
func (l *ll_iter) LLIteratorDelete() (vtype, error) {
	if l.list.IsEmpty() || l.node == nil {
		return nil, errors.New("Invalid for deleting!")
	}

	var ret vtype
	l.list.length--

	if l.LLIteratorHasNext() {

		if l.LLIteratorHasPrev() { // in the middle
			(l.node.pre).next = l.node.next
			(l.node.next).pre = l.node.pre

			ret = l.node.val

		} else { // in the tail
			(l.node.next).pre = nil
			l.list.head = l.node.next

			ret = l.node.val
		}

		l.node = l.node.next

		return ret, nil

	} else if l.LLIteratorHasPrev() { // in the head
		(l.node.pre).next = nil
		l.list.tail = l.node.pre

		ret = l.node.val
		l.node = l.node.pre

		return ret, nil

	} else { // the only one

		ret = l.node.val
		l.node = nil

		return ret, nil
	}

}
