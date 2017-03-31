package fc

import (
	"fmt"
	"testing"
)

func TestCrawler(t *testing.T) {
	path := "testdir"

	doctable, index := CrawlFileTree(path)

	if doctable == nil || index == nil {
		t.Error("CrawlFileTree Failure!")
	}

	fmt.Printf("%v %v\n", doctable.DTNumDocs(), index.MINumWordsInMemIndex())

	index.MIShow()
}
