package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var dirs bool
var files bool
var symlinks bool
var ext string

func init() {
	flag.BoolVar(&dirs, "d", false, "directories")
	flag.BoolVar(&files, "f", false, "files")
	flag.BoolVar(&symlinks, "sl", false, "symlinks")
	flag.StringVar(&ext, "ext", "", "extension" )
	flag.Parse()
	if !dirs && !files && !symlinks && ext == "" {
		dirs = true
		files = true
		symlinks = true
	}
	if len(flag.Args()) != 1 {
		fmt.Println("Wrong argumemts.")
		os.Exit(1)
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if fmt.Sprintf("%s", info.Mode())[1] == 'r' {
		if dirs && info.IsDir() {
			fmt.Println(path)
		}
		if files && info.Mode().IsRegular() {
			if ext == "" || ext != "" && filepath.Ext(path) == "." + ext {
				fmt.Println(path)
			}
		}
		if symlinks && info.Mode().Type() == os.ModeSymlink {
			realPath, err := filepath.EvalSymlinks(path)
			if err != nil {
				fmt.Println(path, "-> [broken]")
			} else {
				fmt.Println(path, "->", realPath)
			}
		}
	}
	return nil
}
func main() {
	input, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println("Bad path: ", err)
	}
	err = input.Close()
	if err != nil {
		return
	}
	err = filepath.Walk(flag.Arg(0), walkFunc)
	if err != nil {
		fmt.Print("Error: ", err)
	}
}