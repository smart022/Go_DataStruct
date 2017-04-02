package main

import (
	_ "../datastructs/linkedlist"
	FC "./crawler"
	_ "./doctable"
	MI "./memindex"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		Usage()
		return
	}

	doctable, index := FC.CrawlFileTree(os.Args[1])

	if doctable == nil || index == nil {
		fmt.Println("Crawl failure!")
		return
	}

	//var input_str string

	inputReader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("Enter query:")
		input_str, err := inputReader.ReadString('\n')
		if len(input_str) == 0 {
			continue
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		query := strings.Fields(input_str)
		if len(query) == 0 {
			continue
		}
		if query[0] == "q" || query[0] == "quit" {
			fmt.Println("Exit!")
			break
		}

		retList := index.MIProcessQuery(query)
		if retList == nil {
			fmt.Println("Not found!")
			continue
		}

		llit, _ := retList.LLMakeIterator()
		count := retList.Len()
		for i := uint32(0); i < count; i++ {
			payload, _ := llit.LLIteratorGetPayload()
			ret, _ := payload.(*(MI.SearchResult))

			fmt.Printf("    %s (%v)\n", doctable.DTLookupDocID(ret.SRGetID()), ret.SRGetRank())
			llit.LLIteratorNext()
		}

	}

}

func Usage() {
	fmt.Printf("Usage: ./searchshell <docroot>\n")
	fmt.Printf("where <docroot> is an absolute or relative ",
		"path to a directory to build an index under.\n")
	fmt.Println("type 'q' or 'quit' to terminate this program in query!\n")
	return
}
