package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

var l bool
var m bool
var w bool

func init() {
	flag.BoolVar(&l, "l", false, "count lines")
	flag.BoolVar(&m, "m", false, "count chars")
	flag.BoolVar(&l, "w", false, "count words")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Error: no arguments")
		os.Exit(1)
	}
	if (l&&m&&w) || (l&&m) || (l&&w) || (m&&w) {
		fmt.Println("Error: use only one flag")
		os.Exit(1)
	}
	if !l && !m && !w {
		w = true
	}
}

func countLines(fileName string, waitGroup *sync.WaitGroup) {
	content, err := os.ReadFile(fileName)
	defer waitGroup.Done()
	if err != nil {
		fmt.Printf("\"%s\" - %v\n", fileName, err)
		return
	}
	count := bytes.Count(content, []byte("\n")) + 1
	fmt.Printf("%d\t%s\n", count, fileName)
}

func countChars(fileName string, waitGroup *sync.WaitGroup) {
	content, err := os.ReadFile(fileName)
	defer waitGroup.Done()
	if err != nil {
		fmt.Printf("\"%s\" - %v\n", fileName, err)
		return
	}
	count := utf8.RuneCount(content)
	fmt.Printf("%d\t%s\n", count, fileName)
}

func countWords(fileName string, waitGroup *sync.WaitGroup) {
	content, err := os.ReadFile(fileName)
	defer waitGroup.Done()
	if err != nil {
		fmt.Printf("\"%s\" - %v\n", fileName, err)
		return
	}
	words := strings.Fields(string(content))
	fmt.Printf("%d\t%s\n", len(words), fileName)
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(flag.Args()))
	if l {
		for _,v := range flag.Args() {
			go countLines(v, &waitGroup)
		}
	}
	if m {
		for _,v := range flag.Args() {
			go countChars(v, &waitGroup)
		}
	}
	if w {
		for _,v := range flag.Args() {
			go countWords(v, &waitGroup)
		}
	}
	waitGroup.Wait()
}