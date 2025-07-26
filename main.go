package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Usage is jlox [script]")
	} else {
		runFile(os.Args[1])
	}
}

func runFile(file string) {
	fmt.Println(file)
}
