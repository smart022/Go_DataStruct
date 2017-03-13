package linkedlist

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLinkedlist(t *testing.T) {
	l := AllocateLinkedList()
	var Num uint32 = 10

	for i := uint32(0); i < Num; i++ {
		l.Append(rand.Uint32())
	}

	if l.Len() != Num {
		t.Error("Append error!")
	}

	cur_node := l.head

	for i := uint32(0); i < Num; i++ {
		fmt.Printf("%v %T:%v\n", i, cur_node.val, cur_node.val)
		cur_node = cur_node.next
	}

	// interface{} type convert assert!!
	l.SortLinkedList(true, func(a, b vtype) bool {
		aval, a_ok := a.(uint32)
		bval, b_ok := b.(uint32)
		if a_ok && b_ok {
			return aval > bval
		} else {
			t.Errorf("compare failure!")
			return false
		}
	})

	fmt.Println()
	cur_node = l.head

	for i := uint32(0); i < Num; i++ {
		fmt.Printf("%v %T:%v\n", i, cur_node.val, cur_node.val)
		cur_node = cur_node.next
	}

	for i := uint32(0); i < Num/2; i++ {
		if _, err := l.Pop(); err != nil {
			t.Errorf("Pop failure0!")
		}

		if l.Len() != Num-i-1 {
			t.Errorf("Pop failure1!")
		}
	}
	for i := uint32(0); i < Num/2; i++ {
		if _, err := l.Slice(); err != nil {
			t.Errorf("Pop failure0!")
		}

		if l.Len() != Num/2-i-1 {
			t.Errorf("Pop failure1!")
		}
	}

	if !l.IsEmpty() {
		t.Errorf("Pop failure!")
	}

}

func TestIterator(t *testing.T) {
	l := AllocateLinkedList()
	var count int = 20
	for i := 0; i < count; i++ {
		l.Append(i)
	}

	iter1, _ := l.LLMakeIterator()

	for i := 0; i < count/2; i++ {

		if val, ok := iter1.LLIteratorGetPayload(); ok != nil || val != i {
			t.Errorf("GetPayload error!!")
		}

		if iter1.LLIteratorHasNext() {
			iter1.LLIteratorNext()
		}

	}

	for i := 0; i < count/2; i++ {

		if val, ok := iter1.LLIteratorGetPayload(); ok != nil || val != count/2-i {
			t.Errorf("GetPayload error!!")
		}

		if iter1.LLIteratorHasPrev() {
			iter1.LLIteratorPrev()
		}

	}

	iter2, _ := l.LLMakeIterator()
	for i := 0; i < count-1; i++ {

		if val, ok := iter2.LLIteratorGetPayload(); ok != nil || val != i {
			t.Errorf("GetPayload error!!")
		}

		iter2.LLIteratorNext()
	}
	if ok := iter2.LLIteratorNext(); ok == nil {
		t.Error("NextMoving error!")
	}

	iter3, _ := l.LLMakeIterator()
	for i := 0; i < count; i++ {

		if val, ok := iter3.LLIteratorDelete(); ok != nil || val != i {
			t.Errorf("LLIteratorDelete error0!!")
		}
		if iter3.list.Len() != uint32(count-i-1) {
			t.Errorf("LLIteratorDelete error1!!")
		}

	}

	if _, ok := iter3.LLIteratorDelete(); ok == nil {
		t.Errorf("LLIteratorDelete error2!!")
	}
}
