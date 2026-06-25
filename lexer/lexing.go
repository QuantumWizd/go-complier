package lexer

type TokenType string

const (
	TOKEN_NUMBER     TokenType = "NUMBER"
	TOKEN_IDENTIFIER TokenType = "IDENTIFIER"
	TOKEN_LET        TokenType = "LET"
	TOKEN_PRINT      TokenType = "PRINT"
	TOKEN_ASSIGN     TokenType = "ASSIGN"
	TOKEN_PLUS       TokenType = "PLUS"
	TOKEN_MINUS      TokenType = "MINUS"
	TOKEN_ASTERISK   TokenType = "ASTERISK"
	TOKEN_SLASH      TokenType = "SLASH"
	TOKEN_EOF        TokenType = "EOF"
	TOKEN_ILLEGAL    TokenType = "ILLEGAL"
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

type Lexer struct {
	input string
	pos   int
	line  int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
		pos:   0,
		line:  1,
	}
}

func (l *Lexer) Tokenize() []Token {
	var tokens []Token

	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == TOKEN_EOF {
			break
		}
	}

	return tokens
}
