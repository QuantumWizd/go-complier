package lexer

type TokenType string

const (
	TOKEN_NUMBER        TokenType = "NUMBER"
	TOKEN_IDENTIFIER    TokenType = "IDENTIFIER"
	TOKEN_ASSIGN        TokenType = "ASSIGN"
	TOKEN_PLUS          TokenType = "PLUS"
	TOKEN_MINUS         TokenType = "MINUS"
	TOKEN_ASTERISK      TokenType = "ASTERISK"
	TOKEN_SLASH         TokenType = "SLASH"
	TOKEN_EOF           TokenType = "EOF"
	TOKEN_ILLEGAL       TokenType = "ILLEGAL"
	TOKEN_VARIABLE      TokenType = "VARIABLE"
	TOKEN_CONSTANT      TokenType = "CONSTANT"
	TOKEN_DECLARE       TokenType = "DECLARE"
	TOKEN_EQUAL         TokenType = "EQUAL"
	TOKEN_NOT_EQUAL     TokenType = "NOT_EQUAL"
	TOKEN_GREATER       TokenType = "GREATER"
	TOKEN_LESS          TokenType = "LESS"
	TOKEN_IF            TokenType = "IF"
	TOKEN_ELSE          TokenType = "ELSE"
	TOKEN_FOR           TokenType = "FOR"
	TOKEN_FUNCTION      TokenType = "FUNCTION"
	TOKEN_RETURN        TokenType = "RETURN"
	TOKEN_LPAREN        TokenType = "LPAREN"
	TOKEN_RPAREN        TokenType = "RPAREN"
	TOKEN_LBRACE        TokenType = "LBRACE"
	TOKEN_RBRACE        TokenType = "RBRACE"
	TOKEN_COMMA         TokenType = "COMMA"
	TOKEN_TRUE          TokenType = "TRUE"
	TOKEN_FALSE         TokenType = "FALSE"
	TOKEN_PRINT         TokenType = "PRINT"
	TOKEN_MODULUS       TokenType = "MODULUS"
	TOKEN_PLUS_ASSIGN   TokenType = "PLUS_ASSIGN"
	TOKEN_MINUS_ASSIGN  TokenType = "MINUS_ASSIGN"
	TOKEN_MUL_ASSIGN    TokenType = "MUL_ASSIGN"
	TOKEN_DIV_ASSIGN    TokenType = "DIV_ASSIGN"
	TOKEN_AND           TokenType = "AND"
	TOKEN_OR            TokenType = "OR"
	TOKEN_NOT           TokenType = "NOT"
	TOKEN_BITWISE_AND   TokenType = "BITWISE_AND"
	TOKEN_BITWISE_OR    TokenType = "BITWISE_OR"
	TOKEN_BITWISE_XOR   TokenType = "BITWISE_XOR"
	TOKEN_LEFT_SHIFT    TokenType = "LEFT_SHIFT"
	TOKEN_RIGHT_SHIFT   TokenType = "RIGHT_SHIFT"
	TOKEN_INCREMENT     TokenType = "INCREMENT"
	TOKEN_DECREMENT     TokenType = "DECREMENT"
	TOKEN_GREATER_EQUAL TokenType = "GREATER_EQUAL" 
	TOKEN_LESS_EQUAL    TokenType = "LESS_EQUAL"    
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
