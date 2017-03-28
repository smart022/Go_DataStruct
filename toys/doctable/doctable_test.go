package dt

import (
	"fmt"
	"testing"
)

func TestDT(t *testing.T) {
	test_arr := []string{
		"California dream",
		"All the leaves are brown",
		"on Such a winter day",
		"when i get down on my knee",
		"ans the sky is grey",
		"Ive been for a while",
	}

	dt := AllocateDocTable()
	if dt == nil {
		t.Error("AllocateDocTable failure!")
	}

	for i, val := range test_arr {
		dt.DTRegisterDocName(val)
		if dt.DTNumDocs() != uint32(i+1) {
			t.Error("DTRegisterDocName failure0!")
		}
	}

	var id uint32
	var idcot []uint32
	for _, val := range test_arr {
		id = dt.DTLookupDocName(val)
		if id == 0 {
			t.Error("DTLookupDocName failure0!")
		}

		idcot = append(idcot, id)
	}

	for _, val := range idcot {

		docname := dt.DTLookupDocID(val)
		if docname == "" {
			t.Error("DTLookupDocID failure0!")
		}

		fmt.Printf("%v:%s\n", val, docname)
	}

}
