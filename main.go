package main

import (
	"interpreters/lox"
	"os"
)

func main() {
	argLen := len(os.Args)
	lox := lox.Lox{}
	if argLen == 1 {
		lox.RunPrompt(os.Stdin)
	} else if argLen == 2 {
		lox.RunFile(os.Args[1])
	} else {
		panic("Usage is jlox [script]")
	}
}
