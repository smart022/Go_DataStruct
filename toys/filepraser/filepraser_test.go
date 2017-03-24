package fp

import (
	HT "../../datastructs/hashtable"

	_ "bytes"
	_ "errors"
	"fmt"
	_ "io"
	_ "io/ioutil"
	_ "os"
	"testing"
	_ "unicode"
)

func TestFP(t *testing.T) {
	in := "filepraser.go" // seems work incorrectly in relative path name
	ht, err := BuildWordHT(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	ht_iter := ht.HTMakeIterator()
	var kvptr *HT.HTKeyValue

	for !ht_iter.HTIteratorPostEnd() {
		kvptr = ht_iter.HTIteratorGet()
		actval, ok := (kvptr.HTKeyValueGet()).(*WordTimes)
		if !ok {
			fmt.Println("failure0!")
			return
		}
		fmt.Printf("%s : %v times\n", actval.word, actval.times)

		ht_iter.HTIteratorNext()
	}

}
