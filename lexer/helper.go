package lexer

import "fmt"

// This function used to get the current character from the input string based on the current position of the lexer. If the position is beyond the length of the input, it returns 0 (indicating end of input).
func (l *Lexer) currentChar() byte {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

// This function used to advance the position of the lexer by one character. If the current character is a newline, it increments the line counter as well.
func (l *Lexer) advance() {
	if l.currentChar() == '\n' {
		l.line++
	}
	l.pos++
}

// This Function used to skip over any whitespace characters (spaces, tabs, newlines, carriage returns) in the input string. It continues advancing the position of the lexer until it encounters a non-whitespace character.
func (l *Lexer) skipWhitespace() {
	for l.currentChar() == ' ' || l.currentChar() == '\t' || l.currentChar() == '\n' || l.currentChar() == '\r' {
		l.advance()
	}
}

// This function used to look up a given word and determine if it is a keyword (like "let" or "print") or an identifier. It returns the appropriate TokenType based on the input word.
func lookupKeyword(word string) TokenType {
	if word == "let" {
		return TOKEN_LET
	}

	if word == "print" {
		return TOKEN_PRINT
	}

	return TOKEN_IDENTIFIER

}

// This function used to read a number from the input string. It starts at the current position of the lexer and continues advancing as long as it encounters digit characters. Once it reaches a non-digit character, it returns the substring representing the number.
func (l *Lexer) readNumber() string {
	start := l.pos
	for isDigit(l.currentChar()) {
		l.advance()
	}
	return l.input[start:l.pos]
}

// This function checks if a given character is a digit (0-9). It returns true if the character is a digit, and false otherwise.
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// This function reads an identifier from the input string. It starts at the current position of the lexer and continues advancing as long as it encounters letter characters (including underscores). Once it reaches a non-letter character, it returns the substring representing the identifier.
func (l *Lexer) readIdentifier() string {
	start := l.pos
	for isLetter(l.currentChar()) {
		l.advance()
	}
	return l.input[start:l.pos]
}

// This function checks if a given character is a letter (a-z, A-Z) or an underscore (_). It returns true if the character is a letter or underscore, and false otherwise.
func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

// This function retrieves the next token from the input string. It first skips any whitespace, then checks the current character to determine what type of token to create. It handles single-character tokens (like '=', '+', '-', '*', '/'), numbers, identifiers, and keywords. If it encounters an unrecognized character, it returns an ILLEGAL token.
func (l *Lexer) NextToken() Token {

	l.skipWhitespace()

	if l.pos >= len(l.input) {
		return Token{Type: TOKEN_EOF, Literal: "", Line: l.line}
	}

	switch l.currentChar() {
	case '=':
		l.advance()
		return Token{Type: TOKEN_ASSIGN, Literal: "=", Line: l.line}
	case '+':
		l.advance()
		return Token{Type: TOKEN_PLUS, Literal: "+", Line: l.line}
	case '-':
		l.advance()
		return Token{Type: TOKEN_MINUS, Literal: "-", Line: l.line}
	case '*':
		l.advance()
		return Token{Type: TOKEN_ASTERISK, Literal: "*", Line: l.line}
	case '/':
		l.advance()
		return Token{Type: TOKEN_SLASH, Literal: "/", Line: l.line}
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
