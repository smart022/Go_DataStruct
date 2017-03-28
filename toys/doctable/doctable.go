package dt

import (
	HT "../../datastructs/hashtable"
)

/*
*  	DocTable: cotains two HashTables
*
*   docID 2 docName                        docName 2 docID
*  ----------------                         -----------------------------------------------
*  | Key  |  Val 							| Key                          |     Val
*  -----------------                        ----------------------------------------------
*  |  4   | "test_tree/README.MD"           | FNV32("test_tree/README.MD") |      4
*  ----------------                         ------------------------------------------------
*  |  1   | "test_tree/bashrc"              | FNV32("test_tree/bashrc")    |      1
*  --------------------                     ----------------------------------------------
*   .
*   .
*   .
*
 */

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

func (d *DocTable) DTLookupDocName(docname string) uint32 {
	res := HT.FNVHash32(docname)
	found, kv := (d.docname2docid).LookupHT(res)
	if !found {
		return 0
	}

	actval, _ := (kv.HTKeyValueGet()).(uint32)

	return actval
}

func (d *DocTable) DTLookupDocID(docid uint32) string {
	found, kv := (d.docid2docname).LookupHT(docid)
	if !found {
		return ""
	}
	actval, _ := (kv.HTKeyValueGet()).(string)

	return actval
}

func (d *DocTable) DTGetDocidTable() *HT.HashTable {
	return d.docid2docname
}
