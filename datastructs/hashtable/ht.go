package ht

import (
	ll "../linkedlist"
	_ "errors"
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
	buckets      []*ll.LinkedList
}

type HTKeyValue struct {
	key uint32
	val interface{}
}

func CreateStringKV(str string) *HTKeyValue {

	ret := new(HTKeyValue)
	ret.key = FNVHash32(str)
	ret.val = str

	return ret
}

func AllocateHashTable(num_buckets uint32) *HashTable {
	if num_buckets == 0 {
		return nil
	}

	ht := new(HashTable)
	ht.num_buckets = num_buckets
	ht.num_elements = 0
	ht.buckets = make([]*ll.LinkedList, num_buckets)

	for i := uint32(0); i < num_buckets; i++ {
		ht.buckets[i] = ll.AllocateLinkedList()
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
func (h *HashTable) InsertHT(kv *HTKeyValue) (bool, *HTKeyValue) {
	if h == nil {
		return false, nil
	}
	// get the list for insertion
	insertBk := h.HashKeyToBucketNum(kv.key)
	insertChain := h.buckets[insertBk]

	// situation 3
	// 1. empty
	if insertChain.Len() == 0 {
		insertChain.Push(kv)

		h.num_elements++

		return true, nil
	}

	lliter, _ := insertChain.LLMakeIterator()

	// 2. same key collison
	if right, oldkv := BucketHasKey(lliter, kv.key); right == true {

		lliter.LLIteratorDelete()
		insertChain.Append(kv)
		h.num_elements++

		return true, oldkv
	} else { // 3. insert literally
		insertChain.Append(kv)
		h.num_elements++

		return true, nil
	}

}

func (h *HashTable) LookupHT(key uint32) (bool, *HTKeyValue) {
	if h == nil {
		return false, nil
	}

	lookupBk := h.HashKeyToBucketNum(key)
	lookupchain := h.buckets[lookupBk]

	if lookupchain.Len() == 0 {
		return false, nil
	}

	lliter, _ := lookupchain.LLMakeIterator()
	//var retKv *HTKeyValue
	if right, retKv := BucketHasKey(lliter, key); right == true {
		return true, retKv
	}

	return false, nil
}

func (h *HashTable) RemoveFromHT(key uint32) (bool, *HTKeyValue) {
	if h == nil {
		return false, nil
	}

	rmBk := h.HashKeyToBucketNum(key)
	rmChain := h.buckets[rmBk]

	if rmChain.IsEmpty() {
		return false, nil
	}

	lliter, _ := rmChain.LLMakeIterator()
	if right, retKv := BucketHasKey(lliter, key); right == true {

		lliter.LLIteratorDelete()
		h.num_elements--

		return true, retKv
	}

	return false, nil
}

func BucketHasKey(iter *ll.LLiter, key uint32) (bool, *HTKeyValue) {
	for {
		kv, _ := iter.LLIteratorGetPayload()
		actkv, kv_ok := kv.(*HTKeyValue)
		if kv_ok && actkv.key == key {
			return true, actkv
		}

		if !iter.LLIteratorHasNext() {
			break
		}

		iter.LLIteratorNext()
	}

	return false, nil
}
