package main

import (
	"flag"
	"fmt"
	"path"
)

func main() {
	fileName := flag.String("f", "", "input file name")
	flag.Parse()
	var dbreader DBReader

	if path.Ext(*fileName) == ".xml" {
		dbreader = XMLname(*fileName)
		content, err:= dbreader.read()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(dbreader.convert(content))
	} else if path.Ext(*fileName) == ".json" {
		dbreader = JSONname(*fileName)
		content, err := dbreader.read()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(dbreader.convert(content))
	} else {
		fmt.Println("usage: ./DBReader -f <filename>.xml")
		fmt.Println("       ./DBReader -f <filename>.json")
	}
}
