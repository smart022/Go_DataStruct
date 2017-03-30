package mdx

import (
	HT "../../datastructs/hashtable"
	_ "../../datastructs/linkedlist"
	"fmt"
	"testing"
)

type Tstr struct {
	word  string
	docid uint32
	times uint32
}

func TestMDX(t *testing.T) {
	mi := AllocateMemIndex()

	if mi.MINumWordsInMemIndex() != 0 {
		t.Error("MINumWordsInMemIndex failure!")
	}

	test_str := []string{
		"chao",
		"chao",
		"chao",
		"hao",
		"hao",
		"hao",
		"liu",
		"chi",
		"ao",
		"pao",
	}

	cot := make([]Tstr, len(test_str))

	var j uint32
	for i, val := range test_str {
		j = uint32(i)
		cot[i] = Tstr{val, j + 1, 8 * (j + 1)}
		ok := mi.MIADDPostingList(val, j+1, 8*(j+1))
		if !ok {
			t.Error("MIADDPostingList failure!")
		}
	}

	query := []string{
		"hao",
	}

	found, kv := ((*HT.HashTable)(mi)).LookupHT(HT.FNVHash32(query[0]))
	if !found {
		t.Error("...")
	}

	wds, _ := (kv.HTKeyValueGet()).(*WordDocSet)
	Hnum := (wds.docIDs).NumElementsOfHashTable()
	fmt.Println((wds.docIDs).NumElementsOfHashTable())

	htiter := (wds.docIDs).HTMakeIterator()

	for i := 0; i < int(Hnum); i++ {
		hkv := htiter.HTIteratorGet()
		fmt.Printf("key:%v\n", hkv.HTKeyValueGetKey())
	}
	// mistake in MIADDPostingL

	///////
	ll := mi.MIProcessQuery(query)

	iter, _ := ll.LLMakeIterator()

	llen := ll.Len()
	for i := uint32(0); i < llen; i++ {
		payload, _ := iter.LLIteratorGetPayload()
		ret, _ := (payload).(*SearchResult)
		fmt.Printf("id:%v rank:%v\n", ret.docid, ret.rank)
	}
}
