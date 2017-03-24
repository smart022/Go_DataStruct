package fp

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

/*
*  this package mainly works for building a HashTable like this
*
*   --------------------------------
*   | key:          | val:         |
*   | FNVHash32(str)| WordTimes *  |
*   --------------------------------
*   | ..            |              |
*   --------------------------------
*    .
*    .
*    .
*
*
*    represent a word distribution of a file
 */

type WordTimes struct {
	word  string
	times uint32
}

// main function: create a new HashTable from a file
func BuildWordHT(filename string) (*HT.HashTable, error) {
	if len(filename) == 0 {
		return nil, errors.New("invalid filename!")
	}

	filebuf, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// without ACSII examing

	tab := HT.AllocateHashTable(64)

	LoopAndInsert(tab, filebuf)

	if tab.NumElementsOfHashTable() == 0 {
		tab = nil
	}

	return tab, nil
}

// static helper functions below
//
func ReadFile(filename string) (*[]byte, error) {
	buf, ioerr := ioutil.ReadFile(filename)
	if ioerr != nil {
		fmt.Println(ioerr)
		return nil, errors.New("ReadFile failure!")
	}

	return &buf, nil

}

func LoopAndInsert(ht *HT.HashTable, filebuf *[]byte) {
	//split
	str := bytes.FieldsFunc(*filebuf, func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	//tolower
	for i := 0; i < len(str); i++ {
		str[i] = bytes.ToLower(str[i])
	}
	mm := make(map[string]uint32)
	for _, val := range str {
		mm[string(val)]++
	}

	for key, val := range mm {
		WTptr := new(WordTimes)
		WTptr.word = key
		WTptr.times = val

		AddToHT(ht, WTptr)
	}
}

func AddToHT(ht *HT.HashTable, WTptr *WordTimes) {
	hashval := HT.FNVHash32(WTptr.word)
	kv := HT.CreateKV(hashval, WTptr)

	ht.InsertHT(kv)

}
