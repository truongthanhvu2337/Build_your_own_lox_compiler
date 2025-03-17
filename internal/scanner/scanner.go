package scanner

import (
	"github.com/codecrafters-io/interpreter-starter-go/internal/errorutil"
	"github.com/codecrafters-io/interpreter-starter-go/internal/token"
)
		
	type Scanner struct {
		source []rune
		tokens []token.Token
		start int
		current int
		line int
	}

	func NewScanner(source string) *Scanner {
		return &Scanner{
			source: []rune(source),
			tokens: []token.Token{}, 
			start: 0, 
			current: 0, 
			line: 1,
		}
	}

	func (s *Scanner) isAtEnd() bool {
		return s.current >= len(s.source)
	}

	//inteface{} is similar to "Object" in Java or C#
	func (s *Scanner) addToken(tokenType token.TokenType, literal interface{}) {
		// String text = source.substring(start, current);
		// tokens.add(new Token(type, text, literal, line));
		text := string(s.source[s.start:s.current])
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

	func (s *Scanner) peek() rune {
		if s.isAtEnd() {
			return '\000'
		}
		return s.source[s.current]
	}

	func (s *Scanner) peekNext() rune {
		if s.current+1 >= len(s.source) {
			return '\000'
		}
		return s.source[s.current+1]
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

	//string literal
	func (s *Scanner) string() {
		for s.peek() != '"' && !s.isAtEnd() {
			if s.peek() == '\n' {
				s.line++
			}
			s.advance()
		}

		if s.isAtEnd() {
			errorutil.ErrorUtil(s.line, "Unterminated string.")
			return
		}

		s.advance()

		value := string(s.source[s.start+1:s.current-1])
		s.addToken(token.STRING, value)
	}

	//number literal
	func (s *Scanner) isDigit(c rune) bool {
		return c >= '0' && c <= '9'
	}

	func (s *Scanner) number() {
		for s.isDigit(s.peek()) {
			s.advance()
		}

		// Look for a fractional part.
		if s.peek() == '.' && s.isDigit(s.peekNext()) {
			s.advance()

			for s.isDigit(s.peek()) {
				s.advance()
			}
		}

		s.addToken(token.NUMBER, string(s.source[s.start:s.current]))
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
		case '!':
			if s.match('=') {
				s.addToken(token.BANG_EQUAL, nil)
			} else {
				s.addToken(token.BANG, nil)
			}
		case '=':
			if s.match('=') {
				s.addToken(token.EQUAL_EQUAL, nil)
			} else {
				s.addToken(token.EQUAL, nil)
			}
		case '<':
			if s.match('=') {
				s.addToken(token.LESS_EQUAL, nil)
			} else {
				s.addToken(token.LESS, nil)
			}
		case '>':
			if s.match('=') {
				s.addToken(token.GREATER_EQUAL, nil)
			} else {
				s.addToken(token.GREATER, nil)
			}
		case '/':
			if s.match('/') { 
				for s.peek() != '\n' && !s.isAtEnd() {
					s.advance() 
				}
			} else {
				s.addToken(token.SLASH, nil) 
			}
		case ' ', '\r', '\t':
			break
		case '\n':
			s.line++
		case '"':
			s.string()
			

		default:
			if s.isDigit(ch) {
				s.number()
			} else {
				errorutil.ErrorUtil(s.line, "Unexpected character: %c", ch)
			}
		}
	}

