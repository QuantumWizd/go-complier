package ast

import "github.com/QuantumWizd/go-complier/lexer"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type NumberLiteral struct {
	Token lexer.Token
	Value string
}

type Identifier struct {
	Token lexer.Token
	Value string
}

type Boolean struct {
	Token lexer.Token
	Value bool
}

type Infix struct {
	Token    lexer.Token
	Right    Expression
	Operator string
	Left     Expression
}

type VariableDeclaration struct {
	Token lexer.Token
	Name  Identifier
	Value Expression
}

type ConstantDeclaration struct {
	Token lexer.Token
	Name  Identifier
	Value Expression
}

type ReturnStatement struct {
	Token lexer.Token
	Value Expression
}

type IfStatement struct {
	Token     lexer.Token
	Condition Expression
	Body      []Statement
	Else      []Statement
}

type ForStatement struct {
	Token lexer.Token
	Body  []Statement
}

type FunctionNode struct {
	Token     lexer.Token
	Name      Identifier
	Parameters []Identifier
	Body      []Statement
}
