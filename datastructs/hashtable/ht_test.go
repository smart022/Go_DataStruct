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
	var num uint32 = 10

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

	PrintIter(iter)

	//for i := 0; i < len(test_str)/2; i++ {
	ok := iter.HTIteratorNext()
	if ok == false {
		t.Error("HTIteratorNext failure0")
		//	break
	}

	//}
	if iter.is_valid {

		PrintIter(iter)
	}
}

func PrintIter(iter *HTIter) {
	payload, _ := iter.bucket_it.LLIteratorGetPayload()
	actval0, ok0 := payload.(*HTKeyValue)
	actval1, ok1 := (actval0.val).(string)
	if ok0 && ok1 {
		fmt.Println(actval1)
	}
}
