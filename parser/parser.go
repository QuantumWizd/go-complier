package parser

import (
	"github.com/QuantumWizd/go-complier/ast"
	"github.com/QuantumWizd/go-complier/lexer"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken lexer.Token
	peekToken    lexer.Token
}

func NewParser(l *lexer.Lexer) *Parser {

	p := &Parser{
		lexer: l,
	}
	p.advance()
	p.advance()
	return p
}

func (p *Parser) Parse() []ast.Statement {

	var statements []ast.Statement

	for p.currentToken.Type != lexer.TOKEN_EOF {
		stmt := p.parseStatement()
		statements = append(statements, stmt)
		p.advance()
	}
	return statements

}

func (p *Parser) parseStatement() ast.Statement {

	switch p.currentToken.Type {
	case lexer.TOKEN_VARIABLE:
		return p.parseVariableDeclaration()
	case lexer.TOKEN_CONSTANT:
		return p.parseConstantDeclaration()
	case lexer.TOKEN_FUNCTION:
		return p.parseFunctionDeclaration()
	case lexer.TOKEN_IF:
		return p.parseIfStatement()
	case lexer.TOKEN_FOR:
		return p.parseForStatement()
	case lexer.TOKEN_RETURN:
		return p.parseReturnStatement()
	default:
		return nil

	}

}

func (p *Parser) parseVariableDeclaration() ast.Statement {
	// step 1 - save variable token
	// token := p.currentToken
	// // step 2 - move to name
	// p.advance()

	// // step 3 - save the name
	// name := ast.Identifier{
	// 	Token: p.currentToken,
	// 	Value: p.currentToken.Literal,
	// }

	return nil

}

func (p *Parser) parseConstantDeclaration() ast.Statement {
	return nil

}

func (p *Parser) parseFunctionDeclaration() ast.Statement {
	return nil

}

func (p *Parser) parseIfStatement() ast.Statement {
	return nil

}

func (p *Parser) parseForStatement() ast.Statement {
	return nil

}

func (p *Parser) parseReturnStatement() ast.Statement {
	return nil

}
