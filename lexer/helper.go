package lexer

import "fmt"

func (l *Lexer) currentChar() byte {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) peekChar() byte {
	if l.pos+1 >= len(l.input) {
		return 0
	}
	return l.input[l.pos+1]
}

func (l *Lexer) advance() {
	if l.currentChar() == '\n' {
		l.line++
	}
	l.pos++
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar() == ' ' || l.currentChar() == '\t' || l.currentChar() == '\n' || l.currentChar() == '\r' {
		l.advance()
	}
}

func lookupKeyword(word string) TokenType {
	switch word {
	case "variable":
		return TOKEN_VARIABLE
	case "constant":
		return TOKEN_CONSTANT
	case "function":
		return TOKEN_FUNCTION
	case "return":
		return TOKEN_RETURN
	case "if":
		return TOKEN_IF
	case "else":
		return TOKEN_ELSE
	case "for":
		return TOKEN_FOR
	case "true":
		return TOKEN_TRUE
	case "false":
		return TOKEN_FALSE
	case "print":
		return TOKEN_PRINT
	default:
		return TOKEN_IDENTIFIER
	}
}

func (l *Lexer) readNumber() string {
	start := l.pos
	for isDigit(l.currentChar()) {
		l.advance()
	}
	return l.input[start:l.pos]
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	start := l.pos
	for isLetter(l.currentChar()) {
		l.advance()
	}
	return l.input[start:l.pos]
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	if l.pos >= len(l.input) {
		return Token{Type: TOKEN_EOF, Literal: "", Line: l.line}
	}

	switch l.currentChar() {
	case '=':
		if l.peekChar() == '=' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_EQUAL, Literal: "==", Line: l.line}
		}
		l.advance()
		return Token{Type: TOKEN_ASSIGN, Literal: "=", Line: l.line}
	case '+':
		if l.peekChar() == '=' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_PLUS_ASSIGN, Literal: "+=", Line: l.line}
		}
		if l.peekChar() == '+' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_INCREMENT, Literal: "++", Line: l.line}
		}
		l.advance()
		return Token{Type: TOKEN_PLUS, Literal: "+", Line: l.line}
	case '-':
		if l.peekChar() == '=' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_MINUS_ASSIGN, Literal: "-=", Line: l.line}
		}
		if l.peekChar() == '-' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_DECREMENT, Literal: "--", Line: l.line}
		}
		l.advance()
		return Token{Type: TOKEN_MINUS, Literal: "-", Line: l.line}
	case '*':
		if l.peekChar() == '=' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_MUL_ASSIGN, Literal: "*=", Line: l.line}
		}
		l.advance()
		return Token{Type: TOKEN_ASTERISK, Literal: "*", Line: l.line}
	case '/':
		if l.peekChar() == '=' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_DIV_ASSIGN, Literal: "/=", Line: l.line}
		}
		l.advance()
		return Token{Type: TOKEN_SLASH, Literal: "/", Line: l.line}
	case '%':
		l.advance()
		return Token{Type: TOKEN_MODULUS, Literal: "%", Line: l.line}
	case '!':
		if l.peekChar() == '=' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_NOT_EQUAL, Literal: "!=", Line: l.line}
		}
		l.advance()
		return Token{Type: TOKEN_NOT, Literal: "!", Line: l.line}
	case '&':
		if l.peekChar() == '&' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_AND, Literal: "&&", Line: l.line}
		}
		l.advance()
		return Token{Type: TOKEN_BITWISE_AND, Literal: "&", Line: l.line}
	case '|':
		if l.peekChar() == '|' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_OR, Literal: "||", Line: l.line}
		}
		l.advance()
		return Token{Type: TOKEN_BITWISE_OR, Literal: "|", Line: l.line}
	case '^':
		l.advance()
		return Token{Type: TOKEN_BITWISE_XOR, Literal: "^", Line: l.line}
	case '<':
		if l.peekChar() == '=' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_LESS_EQUAL, Literal: "<=", Line: l.line}
		}
		if l.peekChar() == '<' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_LEFT_SHIFT, Literal: "<<", Line: l.line}
		}
		l.advance()
		return Token{Type: TOKEN_LESS, Literal: "<", Line: l.line}
	case '>':
		if l.peekChar() == '=' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_GREATER_EQUAL, Literal: ">=", Line: l.line}
		}
		if l.peekChar() == '>' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_RIGHT_SHIFT, Literal: ">>", Line: l.line}
		}
		l.advance()
		return Token{Type: TOKEN_GREATER, Literal: ">", Line: l.line}
	case ':':
		if l.peekChar() == ':' {
			l.advance()
			l.advance()
			return Token{Type: TOKEN_DECLARE, Literal: "::", Line: l.line}
		}
		l.advance()
		return Token{Type: TOKEN_ILLEGAL, Literal: ":", Line: l.line}
	case '(':
		l.advance()
		return Token{Type: TOKEN_LPAREN, Literal: "(", Line: l.line}
	case ')':
		l.advance()
		return Token{Type: TOKEN_RPAREN, Literal: ")", Line: l.line}
	case '{':
		l.advance()
		return Token{Type: TOKEN_LBRACE, Literal: "{", Line: l.line}
	case '}':
		l.advance()
		return Token{Type: TOKEN_RBRACE, Literal: "}", Line: l.line}
	case ',':
		l.advance()
		return Token{Type: TOKEN_COMMA, Literal: ",", Line: l.line}
	default:
		if isDigit(l.currentChar()) {
			num := l.readNumber()
			return Token{Type: TOKEN_NUMBER, Literal: num, Line: l.line}
		}
		if isLetter(l.currentChar()) {
			word := l.readIdentifier()
			tokenType := lookupKeyword(word)
			return Token{Type: tokenType, Literal: word, Line: l.line}
		}
		ch := l.currentChar()
		l.advance()
		return Token{Type: TOKEN_ILLEGAL, Literal: string(ch), Line: l.line}
	}
}

func (t Token) String() string {
	return fmt.Sprintf("Type: %-12s Literal: %-15q Line: %d", t.Type, t.Literal, t.Line)
}
