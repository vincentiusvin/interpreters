package lox

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Lox struct {
	hadError bool
}

func (l *Lox) RunFile(fileName string) error {
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

func (l *Lox) RunPrompt(r io.Reader) error {
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
