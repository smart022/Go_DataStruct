package dt

import (
	HT "../../datastructs/hashtable"
	_ "bufio"
	"bytes"
	"errors"
	"fmt"
	_ "io"
	"io/ioutil"
	_ "os"
	"unicode"
)

type DocTable struct {
	docid2docname *HT.HashTable
	docname2docid *HT.HashTable
	max_id        uint32
}

func AllocateDocTable() *DocTable {
	dt := new(DocTable)
	dt.docid2docname = HT.AllocateHashTable(1024)
	dt.docname2docid = HT.AllocateHashTable(1024)
	dt.max_id = 0

	return dt
}

func (d *DocTable) DTNumDocs() uint32 {
	return (d.docid2docname).NumElementsOfHashTable()
}

func (d *DocTable) DTRegisterDocName(docname string) uint32 {

	var docid uint32

	// check whether it exists or not
	res := HT.FNVHash32(docname)
	if found, val := (d.docname2docid).LookupHT(res); found {
		actval, _ := (val.HTKeyValueGet()).(uint32)

		return actval
	}

	// update the maxid and create the new docid
	d.max_id++
	docid = d.max_id

	//
	kv := HT.CreateKV(docid, docname)
	(d.docid2docname).InsertHT(kv)

	//
	kv = HT.CreateKV(res, docid)
	(d.docname2docid).InsertHT(kv)

	return docid
}
