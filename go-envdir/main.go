package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go-envdir /path/to/dir program")
		return
	}

	env, err := ReadDir(os.Args[1])

	if err != nil {
		os.Exit(111)
	}

	res := RunCmd(os.Args[2:], env)
	os.Exit(res)
}