package linkedlist

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLinkedlist(t *testing.T) {
	l := AllocateLinkedList()

	var Num uint32 = 100

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
