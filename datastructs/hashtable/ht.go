package ht

import (
	"../linkedlist/linkedlist"
	"errors"
	"hash/fnv"
)

/*
*  structure figure
*
*  HashTable
* --------------------------------
* | uint32: num of the buckets   	|
* |------------------------------	|
* | uint32: num of the elements  	|
* |------------------------------	|
* | LinkerList*[] : refer to a 	    |
  | slice of pointers to LinkedList |
* ------------------------------------
*     |				LinkedList* []				     LinkedList
*	  |			---------------------              ----------------------
*     --------> |  [0]: 			| -----------> | uint32 : length   |                   Node
*				|--------------------			   ----------------------				  --------------------
*				|  [1]:				|			   | node* : head      | ---------------->| HTKeyValue* : val |--------
*				|-------------------			   ----------------------				  -------------------
*												   | node* : tail                         |	Node*: prev & next
*
*
*/
type HashTable struct {
	num_buckets  uint32
	num_elements uint32
	buckets      []*LinkedList
}

type HTKeyValue struct {
	key uint32
	val interface{}
}

func AllocateHashTable(uint32 num_buckets) *HashTable {
	if num_buckets == 0 {
		return nil
	}

	ht := new(HashTable)
	ht.num_buckets = num_buckets
	ht.num_elements = 0
	ht.buckets = make([]*LinkedList, num_buckets)

	for i := num_buckets - 1; i >= 0; i-- {
		ht.buckets[i] = AllocateLinkedList()
	}

	return ht
}

func (t *HashTable) NumElementsOfHashTable() uint32       { return t.num_elements }
func (t *HashTable) HashKeyToBucketNum(key uint32) uint32 { return key % t.num_buckets }

// usage of fnv
func FNVHash32(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}

func FNVHashInt32(hashme uint32) uint32 {
	h := fnv.New32a()
	detail := make([]byte, 4)
	for i := uint8(0); i < 4; i++ {
		detail[i] = byte((hashme >> (i * 8)) & 0xff)
	}
	h.Write(detail)
	return h.Sum32()
}

// insert success return true
// HTKeyValue return val when there is a same hash key
// and alternate it with this inserted one
func (h *HashTable) Insert(kv HTKeyValue) (bool, *HTKeyValue) {
	if h == nil {
		return false, nil
	}
	// get the list for insertion
	insertBk := h.HashKeyToBucketNum(kv.key)
	insertChain := h.buckets[insertBk]

	// situation 3
	// 1. empty
	if insertChain.Len() == 0 {
		insertChain.Push(&kv)

		h.num_elements++

		return true, nil
	}

	lliter, _ := insertChain.LLMakeIterator()

	// 2. same key collison
	if right, oldkv := BucketHasKey(ll_iter, kv.key); right == true {

		ll_iter.LLIteratorDelete()
		insertChain.Append(&kv)
		h.num_elements++

		return true, oldkv
	} else { // 3. insert literally
		insertChain.Append(&kv)
		h.num_elements++

		return true, nil
	}

}

func BucketHasKey(iter ll_iter, key uint32) (bool, *HTKeyValue) {
	for {
		kv, _ := iter.LLIteratorGetPayload()
		if kv.key == key {
			return true, kv
		}

		if !ll_iter.LLIteratorHasNext() {
			break
		}

		ll_iter.LLIteratorNext()
	}

	return false, nil
}
