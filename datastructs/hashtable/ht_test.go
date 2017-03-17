package ht

import (
	"fmt"
	"testing"
)

func TestHT(t *testing.T) {
	var num uint32 = 100

	ht := AllocateHashTable(num)

	test_str := []string{
		"dnoaskhdl",
		"qkwhedljkasld",
		"ioqwuheolkjasdn",
		"I know who u r",
		"u such bitch",
		"get away from here!",
		"dont approching!",
		"wtf a u doing",
		"uhhhhhh!!!##$$",
		"dragon offsprings",
	}

	if ht.num_elements != 0 || ht.num_buckets != num || ht == nil {
		t.Error("AllocateHashTable failure0")
	}

	for _, val := range test_str {

		ht.InsertHT(CreateStringKV(val))
	}

	if ht.num_elements != uint32(len(test_str)) {
		t.Error("InsertHT failure0!")
	}

	for _, val := range test_str {

		right, kv := ht.LookupHT(FNVHash32(val))

		if !right {
			t.Error("LookupHT failure0!")

		}

		if actval, ok := (kv.val).(string); ok {
			if actval != val {
				t.Error("LookupHT failure1!")
			}
		}

	}

	for i, val := range test_str {
		right, kv := ht.RemoveFromHT(FNVHash32(val))
		if !right {
			t.Error("RemoveFromHT failure0!!")

		}
		if actval, ok := (kv.val).(string); ok {
			if actval != val {
				t.Error("RemoveFromHT failure1!")
			}
		}

		if ht.num_elements != uint32(len(test_str)-i-1) {
			t.Error("RemoveFromHT failure2!!")
		}

	}

}

func TestHTIterator(t *testing.T) {
	var num uint32 = 8

	ht := AllocateHashTable(num)

	test_str := []string{
		"dnoaskhdl",
		"qkwhedljkasld",
		"ioqwuheolkjasdn",
		"I know who u r",
		"u such bitch",
		"get away from here!",
		"dont approching!",
		"wtf a u doing",
		"uhhhhhh!!!##$$",
		"dragon offsprings",
		"Smart is gagaga!!",
		"WowWowWow!!!",
	}

	for _, val := range test_str {

		ht.InsertHT(CreateStringKV(val))
	}

	iter := ht.HTMakeIterator()

	if iter.is_valid == false {
		t.Error("HTMakeIterator failure0!!")
	}

	for a := int(num - 1); a > -1; a-- {
		fmt.Printf("%v : %v\n", a, iter.ht.buckets[a].Len())
	}
	fmt.Printf("sum: %v\n", len(test_str))

	PrintIter(iter)

	for i := 0; i < len(test_str)-1; i++ {

		ok := iter.HTIteratorNext()
		if ok == false {
			t.Error("HTIteratorNext failure0")
			//	break
		}
		Kv := iter.HTIteratorGet()
		PrintKv(Kv)
		PrintIter(iter)

	}

	if iter.bucket_it == nil {
		t.Error("HTIteratorNext failure1")
	}

	if ok := iter.HTIteratorNext(); ok == true {
		t.Error("HTIteratorNext failure2")

	}
	if iter.is_valid {
		t.Error("HTIteratorNext failure3")
	}

	iter1 := ht.HTMakeIterator()

	for i := 0; i < len(test_str); i++ {
		ok, Kv := iter1.HTIteratorDelete()
		if !ok {
			t.Error("HTIteratorDelete failure0")
		}

		if Kv == nil {
			t.Error("HTIteratorDelete failure1")
		}

		PrintKv(Kv)
	}

	if iter1.is_valid {
		t.Error("HTIteratorDelete failure2")
	}

}

func PrintIter(iter *HTIter) {
	payload, _ := iter.bucket_it.LLIteratorGetPayload()
	actval0, ok0 := payload.(*HTKeyValue)
	if ok0 {
		PrintKv(actval0)
	}
}

func PrintKv(Kv *HTKeyValue) {
	actval, ok := (Kv.val).(string)
	if ok {
		fmt.Println(actval)
	}
}
