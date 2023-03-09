package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Pass command like 'wc -l' or 'ls -la'")
		os.Exit(1)
	}
	args := os.Args[2:]
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	input := strings.Fields(string(in))
	for _,v := range input {
		cmd := exec.Command(os.Args[1], append(args, v)...)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println(string(stdout))
	}
}