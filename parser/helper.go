package parser

func (p *Parser) advance() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}
