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

type TokenType string

const (
	LEFT_PAREN  TokenType = "("
	RIGHT_PAREN TokenType = ")"
	LEFT_BRACE  TokenType = "{"
	RIGHT_BRACE TokenType = "{"
	COMMA       TokenType = ","
	DOT         TokenType = "."
	MINUS       TokenType = "-"
	PLUS        TokenType = "+"
	SEMICOLON   TokenType = ";"
	SLASH       TokenType = "/"
	STAR        TokenType = "*"

	BANG          TokenType = "!"
	BANG_EQUAL    TokenType = "!="
	EQUAL         TokenType = "="
	EQUAL_EQUAL   TokenType = "=="
	GREATER       TokenType = ">"
	GREATER_EQUAL TokenType = ">="
	LESS          TokenType = "<"
	LESS_EQUAL    TokenType = "<="

	// dynamic types
	IDENTIFIER TokenType = "IDENTIFIER"
	STRING     TokenType = "STRING"
	NUMBER     TokenType = "NUMBER"

	AND   TokenType = "AND"
	CLASS TokenType = "CLASS"
	ELSE  TokenType = "ELSE"
	FALSE TokenType = "FALSE"
	FUN   TokenType = "FUN"
	FOR   TokenType = "FOR"
	IF    TokenType = "IF"
	NIL   TokenType = "NIL"
	OR    TokenType = "OR"

	PRINT  TokenType = "PRINT"
	RETURN TokenType = "RETURN"
	SUPER  TokenType = "SUPER"
	THIS   TokenType = "THIS"
	TRUE   TokenType = "TRUE"
	VAR    TokenType = "VAR"
	WHILE  TokenType = "WHILE"

	EOF TokenType = "EOF"
)
