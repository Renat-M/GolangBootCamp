package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)
func getArchiveName(filename string, path string) string {
	name := ""
	timestamp := ""
	if strings.HasSuffix(filename, ".log") {
		name = strings.TrimSuffix(filename, ".log")
	} else {
		fmt.Println("Incorrect extension: ", filename)
		os.Exit(1)
	}
	timestamp = strconv.Itoa((int(time.Now().Unix())))
	name += "_" + timestamp + ".tar.gz"
	if path != "" {
		name = path + "/" + filepath.Base(name)
	}
	return name
}

func createArchive(filename string, out io.Writer) error {
	tw := tar.NewWriter(out)
	defer tw.Close()
	file,err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	info,err := file.Stat()
	if err != nil {
		return err
	}
	header,err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}
	header.Name = filename
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}
	_,err = io.Copy(tw, file)
	if err != nil {
		return err
	}
	return nil
}

func archive(filename string, path string, wg *sync.WaitGroup) {
	archname := getArchiveName(filename, path)
	out,err := os.Create(archname)
	if err != nil {
		fmt.Println("Error writing archive: ", err)
		os.Exit(1)
	}
	defer out.Close()
	err = createArchive(filename, out)
	if err != nil {
		fmt.Println("Error creating archive: ", err)
		os.Exit(1)
	}
	wg.Done()
}

func main() {
	var a string
	flag.StringVar(&a, "a", "", "archive dir")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Pass one or more filenames to archive them")
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(flag.Args()))
	for _,v := range flag.Args() {
		go archive(v, a, &wg)
	}
	wg.Wait()
}