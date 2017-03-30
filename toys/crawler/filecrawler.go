package fc

import (
	DT "../doctable"
	FP "../filepraser"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// main function, still lack of MemIndex
func CrawlFileTree(rootdir string) (*DocTable, *MemIndex) {
	path := rootdir
	doctable := DT.AllocateDocTable()
	index := AllocateMemIndex()
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
	}
}

// helper func
func HandleFile(path string, dt *DT.Doctable, index *MI.MemIndex) {

	tab, err := FP.BuildWordHT(path)
	if err != nil {
		//deal
		fmt.Println(err)
		return
	}

	docID := dt.DTRegisterDocName(path)

	it := tab.HTMakeIterator()
	for {
		//deal it
	}

}
