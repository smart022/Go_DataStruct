package ht

import (
	ll "../linkedlist"
	_ "errors"
	_ "fmt"
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
func CreateKV(key uint32, val interface{}) *HTKeyValue {

	ret := new(HTKeyValue)
	ret.key = key
	ret.val = val

	return ret
}

func (h *HTKeyValue) HTKeyValueGet() interface{} {
	return h.val
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

// HashTable iterator
//
//
//
//
type HTIter struct {
	is_valid   bool
	ht         *HashTable
	bucket_num uint32
	bucket_it  *ll.LLiter
}

func (h *HashTable) HTMakeIterator() *HTIter {
	iter := new(HTIter)

	if h.num_elements == 0 {
		iter.is_valid = false
		iter.ht = h
		iter.bucket_it = nil
		return iter
	}
	iter.is_valid = true
	iter.ht = h

	for i := uint32(0); i < h.num_buckets; i++ {
		if (h.buckets[i]).Len() > 0 {
			iter.bucket_num = i
			break
		}
	}

	iter.bucket_it, _ = (h.buckets[iter.bucket_num]).LLMakeIterator()

	return iter
}

func (h *HTIter) HTIteratorNext() bool {
	if !h.is_valid {
		return false
	}

	// there is something left in this Bucket
	if (h.bucket_it).LLIteratorHasNext() {
		(h.bucket_it).LLIteratorNext()
		return true
	}

	// otherwise in other Buckets
	var i uint32

	// pay heed to variable's effect domain
	for i = h.bucket_num + 1; i < h.ht.num_buckets; i++ {
		if (h.ht.buckets[i]).Len() > 0 {

			h.bucket_num = i
			//	fmt.Printf("Stop at %v buckets\n", i)
			break
		}
	}

	// and we foud the bucket
	if i < h.ht.num_buckets {
		h.bucket_it = nil
		//fmt.Printf("current %v bucket's lenght==%v\n", pos, (h.ht.buckets[pos]).Len())
		newBkIter, err := (h.ht.buckets[i]).LLMakeIterator()
		if err != nil {
			//fmt.Println(err)
			return false
		}
		h.bucket_it = newBkIter
		//fmt.Printf("MakeIterator!\n")
		return true
	}

	//fmt.Printf("Not found!\n")
	// else none left, alread at the tail
	h.is_valid = false
	h.bucket_it = nil

	return false
}

func (h *HTIter) HTIteratorPostEnd() bool {
	if !h.is_valid {
		return true
	}

	return false

}
func (h *HTIter) HTIteratorGet() *HTKeyValue {
	if !h.is_valid {
		return nil
	}

	payload, err := h.bucket_it.LLIteratorGetPayload()
	if err != nil {
		return nil
	}
	actkv, ok := payload.(*HTKeyValue)
	if !ok {
		return nil
	}

	return actkv
}

// spare effort to keep HTIter valid
func (h *HTIter) HTIteratorDelete() (bool, *HTKeyValue) {
	if !h.is_valid {
		return false, nil
	}

	table := h.ht
	retKv := h.HTIteratorGet()

	h.HTIteratorNext()

	if ok, _ := table.RemoveFromHT(retKv.key); !ok {
		return false, nil
	}

	return true, retKv

}
