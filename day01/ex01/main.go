package main

import (
	"flag"
	"fmt"
	"path"
)

func main() {
	oldDB := flag.String("old", "", "old database")
	newDB := flag.String("new", "", "new database")
	flag.Parse()
	var dbreaderOld DBReader
	var dbreaderNew DBReader

	if path.Ext(*oldDB) == ".xml" && path.Ext(*newDB) == ".json" {
		dbreaderOld = XMLname(*oldDB)
		old, err := dbreaderOld.read()
		if err != nil {
			fmt.Println("old database is broken")
			return
		}
		dbreaderNew = JSONname(*newDB)
		new, err := dbreaderNew.read()
		if err != nil {
			fmt.Println("new database is broken")
			return
		}
		compare(old, new)
	} else if path.Ext(*newDB) == ".xml" && path.Ext(*oldDB) == ".json" {
		dbreaderOld = JSONname(*oldDB)
		old, err := dbreaderOld.read()
		if err != nil {
			fmt.Println("old database is broken")
			return
		}
		dbreaderNew = XMLname(*newDB)
		new, err := dbreaderNew.read()
		if err != nil {
			fmt.Println("new database is broken")
			return
		}
		compare(new, old)
	} else {
		fmt.Println("wrong extension.")
		fmt.Println("usage: ./compareDB --old <filename>.xml --new <filename>.json")
	}
}
