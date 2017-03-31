package fc

import (
	DT "../doctable"
	FP "../filepraser"
	MI "../memindex"
	"errors"
	_ "flag"
	"fmt"
	"os"
	"path/filepath"
)

// main function, still lack of MemIndex
func CrawlFileTree(rootdir string) (*DT.DocTable, *MI.MemIndex) {
	path := rootdir
	doctable := DT.AllocateDocTable()
	index := MI.AllocateMemIndex()
	// walk through the dir
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		// filter what is not file
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		// deal with each file

		HandleFile(path, doctable, index)

		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
		return nil, nil
	}

	return doctable, index
}

// helper func
func HandleFile(path string, dt *DT.DocTable, index *MI.MemIndex) error {

	tab, err := FP.BuildWordHT(path)
	if err != nil {
		//deal
		return err
	}

	docID := dt.DTRegisterDocName(path)

	iter := tab.HTMakeIterator()
	for tab.NumElementsOfHashTable() != 0 {
		//deal it
		ok, kv := iter.HTIteratorDelete()
		if !ok {
			return errors.New("HTIteratorDelete failure!")
		}

		wt := (kv.HTKeyValueGet()).(*FP.WordTimes)

		index.MIADDPostingList(wt.WTGetWord(), docID, wt.WTGetTimes())
	}

	return nil
}
