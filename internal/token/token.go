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

	BANG_EQUAL      TokenType = "BANG_EQUAL"
	EQUAL_EQUAL     TokenType = "EQUAL_EQUAL"
	LESS_EQUAL      TokenType = "LESS_EQUAL"
	GREATER_EQUAL   TokenType = "GREATER_EQUAL"
	BANG         	TokenType = "BANG"
	EQUAL         	TokenType = "EQUAL"
	LESS         	TokenType = "LESS"
	GREATER         TokenType = "GREATER"

	IDENTIFIER TokenType = "IDENTIFIER"
	STRING     TokenType = "STRING"
	NUMBER     TokenType = "NUMBER"

	AND    TokenType = "AND"
	CLASS  TokenType = "CLASS"
	ELSE   TokenType = "ELSE"
	FALSE  TokenType = "FALSE"
	FUN    TokenType = "FUN"
	FOR    TokenType = "FOR"
	IF     TokenType = "IF"
	NIL    TokenType = "NIL"
	OR     TokenType = "OR"
	PRINT  TokenType = "PRINT"
	RETURN TokenType = "RETURN"
	SUPER  TokenType = "SUPER"
	THIS   TokenType = "THIS"
	TRUE   TokenType = "TRUE"
	VAR    TokenType = "VAR"
	WHILE  TokenType = "WHILE"

	EOF         TokenType = "EOF"
)

var Keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}


type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

//similar to ToString() in Java or C#
func (t Token) String() string {
	literalValue := "null"
	if t.Literal != nil {
		switch value := t.Literal.(type) {
			case float64:
				if value == float64(int(value)) {
					literalValue = fmt.Sprintf("%.1f", value) 
				} else {
					literalValue = fmt.Sprintf("%.10g", value) 
				}
			default:
				literalValue = fmt.Sprintf("%v", t.Literal)
		}
	}
	return fmt.Sprintf("%s %s %s", t.Type, t.Lexeme, literalValue)
}
