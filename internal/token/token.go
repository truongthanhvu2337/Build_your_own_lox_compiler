package token

import "fmt"

type TokenType string

const (
	LEFT_PAREN  TokenType = "LEFT_PAREN"
	RIGHT_PAREN TokenType = "RIGHT_PAREN"
	LEFT_BRACE  TokenType = "LEFT_BRACE"
	RIGHT_BRACE TokenType = "RIGHT_BRACE"
	COMMA       TokenType = "COMMA"
	DOT         TokenType = "DOT"
	MINUS       TokenType = "MINUS"
	PLUS        TokenType = "PLUS"
	SEMICOLON   TokenType = "SEMICOLON"
	SLASH       TokenType = "SLASH"
	STAR        TokenType = "STAR"
	EOF         TokenType = "EOF"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func (t Token) String() string {
	literalValue := "null"
	if t.Literal != nil {
		literalValue = fmt.Sprintf("%v", t.Literal)
	}
	return fmt.Sprintf("%s %s %s", t.Type, t.Lexeme, literalValue)
}