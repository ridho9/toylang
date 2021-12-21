package main

import (
	"fmt"
)

func tokenizeFile(content string) ([]Token, error) {
	t := newTokenizer(content)
	return t.tokenize()
}

type tokenizer struct {
	content     string
	clen        int
	cursor      int
	startCursor int
	result      []Token

	curLine int
	curCol  int
	err     error
}

func newTokenizer(content string) *tokenizer {
	result := &tokenizer{content: content, curLine: 1, curCol: 1}
	result.clen = len(content)
	return result
}

func (t *tokenizer) tokenize() ([]Token, error) {
	t.skipWhiteline()
	for t.canRun() && (t.err == nil) {
		t.startCursor = t.cursor

		c := t.readChar()
		switch c {
		case '+':
			t.appendToken(Plus)
		case '-':
			t.appendToken(Minus)
		case '*':
			t.appendToken(Star)
		case '/':
			t.appendToken(FSlash)
		case '(':
			t.appendToken(LParen)
		case ')':
			t.appendToken(RParen)
		case '=':
			t.appendToken(Equal)
		default:
			if charIsAlpha(c) {
				t.consumeIdent()
				t.appendToken(Ident)
			} else if charIsNum(c) {
				t.consumeNumber()
				t.appendToken(Number)
			} else {
				t.emitError(fmt.Sprintf("unknown char '%c'", c))
			}
		}

		t.skipWhiteline()
	}

	return t.result, t.err
}

func (t *tokenizer) consumeIdent() string {
	for {
		c := t.peekChar()
		if charIsAlpha(c) || charIsNum(c) || c == '_' {
			t.readChar()
		} else {
			break
		}
	}
	return t.currentLex()
}

func (t *tokenizer) consumeNumber() string {
	for {
		c := t.peekChar()
		if charIsNum(c) {
			t.readChar()
		} else {
			break
		}
	}
	return t.currentLex()
}

func (t *tokenizer) readChar() byte {
	if !t.canRun() {
		return 0
	}
	c := t.content[t.cursor]
	t.cursor += 1
	return c
}

func (t *tokenizer) peekChar() byte {
	if !t.canRun() {
		return 0
	}
	return t.content[t.cursor]
}

func (t *tokenizer) skipWhiteline() {
	for t.canRun() {
		c := t.peekChar()
		if c == ' ' || c == '\t' || c == '\n' {
			c = t.readChar()
			t.curCol += 1
			if c == '\n' {
				t.startCursor += 1
				t.appendToken(Newline)
				t.curLine += 1
				t.curCol = 1
			}
			t.startCursor = t.cursor
		} else {
			break
		}
	}
}

func (t *tokenizer) currentLex() string {
	return t.content[t.startCursor:t.cursor]
}

func (t *tokenizer) appendToken(typ TokenType) {
	tok := Token{
		typ:  typ,
		lex:  t.currentLex(),
		line: t.curLine,
		col:  t.curCol,
	}
	t.curCol += len(tok.lex)
	t.result = append(t.result, tok)
}

func (t *tokenizer) canRun() bool {
	return (t.cursor < t.clen)
}

func (t *tokenizer) emitError(msg string) {
	t.err = fmt.Errorf("[%d:%d] %s", t.curLine, t.curCol, msg)
}

func charIsAlpha(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func charIsNum(c byte) bool {
	return ('0' <= c && c <= '9')
}
