package scanner

import (
	"github.com/codecrafters-io/interpreter-starter-go/internal/errorutil"
	"github.com/codecrafters-io/interpreter-starter-go/internal/token"
)
	
type Scanner struct {
	source string
	tokens []token.Token
	start int
	current int
	line int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source, 
		tokens: []token.Token{}, 
		start: 0, 
		current: 0, 
		line: 1,
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) addToken(tokenType token.TokenType, literal interface{}) {
	// String text = source.substring(start, current);
    // tokens.add(new Token(type, text, literal, line));
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, token.Token{
		Type:    tokenType,
		Lexeme:  text,
		Literal: literal,
		Line:    s.line,
	})
}

func (s *Scanner) advance() rune {
	r := []rune(s.source)[s.current]
	s.current++ 
	return r
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if []rune(s.source)[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) ScanTokens() []token.Token {
    for !s.isAtEnd() {
        s.start = s.current
        s.scanToken() 
    }

	s.tokens = append(s.tokens, token.Token{
		Type:    token.EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    s.line,
	})

	return s.tokens
}


//Recongnize lexemes
func (s *Scanner) scanToken() {
	ch := s.advance()
	switch ch {
	case '(':
		s.addToken(token.LEFT_PAREN, nil)
	case ')':
		s.addToken(token.RIGHT_PAREN, nil)
	case '{':
		s.addToken(token.LEFT_BRACE, nil)
	case '}':
		s.addToken(token.RIGHT_BRACE, nil)
	case ',':
		s.addToken(token.COMMA, nil)
	case '.':
		s.addToken(token.DOT, nil)
	case '-':
		s.addToken(token.MINUS, nil)
	case '+':
		s.addToken(token.PLUS, nil)
	case ';':
		s.addToken(token.SEMICOLON, nil)
	case '*':
		s.addToken(token.STAR, nil)
	default:
		errorutil.ErrorUtil(s.line, "Unexpected character: %c", ch)
	}
}

