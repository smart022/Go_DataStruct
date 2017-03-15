package ht

import (
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

		if right, _ := ht.LookupHT(FNVHash32(val)); !right {
			t.Error("LookupHT failure0!")
		}
	}

	for i, val := range test_str {

		if right, _ := ht.RemoveFromHT(FNVHash32(val)); !right {
			t.Error("RemoveFromHT failure0!!")
		}

		if ht.num_elements != uint32(len(test_str)-i-1) {
			t.Error("RemoveFromHT failure1!!")
		}

	}

}
