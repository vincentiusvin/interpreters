package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	argLen := len(os.Args)
	lox := Lox{}
	if argLen == 1 {
		lox.runPrompt(os.Stdin)
	} else if argLen == 2 {
		lox.runFile(os.Args[1])
	} else {
		panic("Usage is jlox [script]")
	}
}

type Lox struct {
	hadError bool
}

func (l *Lox) runFile(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	l.run(string(file))
	if l.hadError {
		return fmt.Errorf("code error")
	}
	return nil
}

func (l *Lox) runPrompt(r io.Reader) error {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		l.run(line)
		l.hadError = false
	}

	return nil
}

func (l *Lox) run(line string) {
	fmt.Println(line)
}

func (l *Lox) report(line int, where string, message string) {
	fmt.Printf("[line %v] Error %v: %v", line, where, message)
	l.hadError = true
}
