package main

import "fmt"

type TokenType int64

//go:generate stringer -type=TokenType
const (
	Plus TokenType = iota
	Minus
	Star
	FSlash
	LParen
	RParen
	Equal
	Newline
	Ident
	Number
	EOF
)

type Token struct {
	typ  TokenType
	lex  string
	line int
	col  int
}

func (t Token) String() string {
	return fmt.Sprintf("[%d:%d %v %#v]", t.line, t.col, t.typ, t.lex)
}
