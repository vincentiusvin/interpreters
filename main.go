package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	argLen := len(os.Args)
	if argLen == 1 {
		runPrompt()
	} else if argLen == 2 {
		runFile(os.Args[1])
	} else {
		panic("Usage is jlox [script]")
	}
}

func runFile(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	run(string(file))
	return nil
}

func runPrompt() error {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		run(line)
	}

	return nil
}

func run(line string) {
	fmt.Println(line)
}
