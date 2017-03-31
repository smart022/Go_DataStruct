package mdx

import (
	HT "../../datastructs/hashtable"
	LL "../../datastructs/linkedlist"
	_ "fmt"
)

/*  this package mainly intend to build such a structure: inverted hashtable
*
*                  MemIndex(HashTable)
*     --------------------------------------
*     | Key: word's Hash | Val: *WordDocSet|                       WordDocSet(struct)
*     --------------------------------------            -----------------------------
*     |FNV32("course")   |          ------------------> | word(string) | *HashTable |                docIDs(HashTable)
*     --------------------------------                  -----------------------------            -----------------------------
*     |                                                 | "course"     |      ---------------->  | Key: docID  | Val: times  |
*     ------------------------------                    ---------------------                    ------------------------
*                                                                                                |  4          |  78
*                                                                                                ------------------------
*                                                                                                |  5          |  63
*                                                                                                ------------------------
*
*
*
*
*
*
*	the toughest one !
*
*

 */

type MemIndex HT.HashTable

type WordDocSet struct {
	word   string
	docIDs *HT.HashTable
}

type SearchResult struct {
	docid uint32
	rank  uint32
}

func AllocateMemIndex() *MemIndex {
	mi := HT.AllocateHashTable(128)
	return (*MemIndex)(mi)
}

func (mi *MemIndex) MINumWordsInMemIndex() uint32 {
	return ((*HT.HashTable)(mi)).NumElementsOfHashTable()
}

// Arguments respectly:
// word:   word u wanna insert
// docid:  which docid this word come from
// times:  how many times this word show up in this doc
func (mi *MemIndex) MIADDPostingList(word string, docid uint32, times uint32) bool {
	var found bool
	var kv *HT.HTKeyValue
	var wds *WordDocSet

	wordkey := HT.FNVHash32(word)

	//fmt.Printf("66. word:%s\n", word)
	found, kv = ((*HT.HashTable)(mi)).LookupHT(wordkey)
	// the first time this inverted index has seen this word.
	if !found {
		wds = new(WordDocSet)
		wds.word = word
		wds.docIDs = HT.AllocateHashTable(128)

		wds_kv := HT.CreateKV(docid, times)

		ok, _ := (wds.docIDs).InsertHT(wds_kv)
		if !ok {
			//
			return false
		}

		mi_kv := HT.CreateKV(wordkey, wds)

		ok, _ = ((*HT.HashTable)(mi)).InsertHT(mi_kv)
		if !ok {
			//
			return false
		}

		//	fmt.Printf("91.word:%s first found\n\n", word)

		return true
	} else { // this word has existed
		//	fmt.Printf("91.word:%s not the first time\n", word)
		wds, _ = (kv.HTKeyValueGet()).(*WordDocSet)
	}

	// then we make sure this docid is the first time we see
	found, kv = (wds.docIDs).LookupHT(docid)
	if found {
		return false
	}

	newKv := HT.CreateKV(docid, times)
	//fmt.Printf("106. CreateKV:%v -> %v \n\n", docid, times)
	ok, _ := (wds.docIDs).InsertHT(newKv)
	if !ok {
		return false
	}

	return true
}

// deal with each query
func (mi *MemIndex) MIProcessQuery(query []string) *LL.LinkedList {
	qlen := len(query)
	if qlen == 0 {
		return nil
	}

	retlist := LL.AllocateLinkedList()

	wordkey := HT.FNVHash32(query[0])
	found, kv := ((*HT.HashTable)(mi)).LookupHT(wordkey)
	if !found {
		return nil
	} else {
		wds, _ := (kv.HTKeyValueGet()).(*WordDocSet)

		ele_num := (wds.docIDs).NumElementsOfHashTable()
		iter := (wds.docIDs).HTMakeIterator()

		for j := uint32(0); j < ele_num; j++ {

			ret := new(SearchResult)

			curkv := iter.HTIteratorGet()

			ret.docid = curkv.HTKeyValueGetKey()
			ret.rank, _ = (curkv.HTKeyValueGet()).(uint32)
			// keep the first word's every showing up case in this retlist
			retlist.Append(ret)
			// without sorting

			iter.HTIteratorNext()

		}

	}
	// only one word, so we finished
	if qlen == 1 {
		return retlist
	}
	// else deal withing other left
	for i := 1; i < qlen; i++ {

		wordkey = HT.FNVHash32(query[i])
		found, kv = ((*HT.HashTable)(mi)).LookupHT(wordkey)

		// not found, so not match exist
		if !found {
			return nil
		}

		llit, _ := retlist.LLMakeIterator()
		ne := retlist.Len()
		wds, _ := (kv.HTKeyValueGet()).(*WordDocSet)

		for j := uint32(0); j < ne; j++ {
			payload, _ := llit.LLIteratorGetPayload()

			ret, _ := payload.(*SearchResult)

			// search this word's docIDS hashtable to find whether
			// it has a same distribution with previous words
			found0, kv0 := (wds.docIDs).LookupHT(ret.docid)
			if found0 {
				adding, _ := (kv0.HTKeyValueGet()).(uint32)
				ret.rank += adding

				llit.LLIteratorNext()

			} else {
				llit.LLIteratorDelete()
			}
		}

	}

	if retlist.Len() == 0 {
		return nil
	}

	// We may Sort
	retlist.SortLinkedList(false, func(a, b interface{}) bool {
		acta, _ := a.(*SearchResult)
		actb, _ := b.(*SearchResult)

		return acta.rank > actb.rank
	})

	return retlist
}
