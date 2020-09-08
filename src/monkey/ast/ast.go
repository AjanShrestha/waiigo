package ast

import "monkey/token"

// Node is the base Node interface
type Node interface {
	TokenLiteral() string
}

// Statement nodes implement this interface
type Statement interface {
	Node
	statementNode()
}

// Expression nodes implement this interface
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the token associated with the node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Identifier node
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral for Identifiers
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// LetStatement is the node for let statement
type LetStatement struct {
	Token token.Token // the token.Let Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral for LetStatement
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// ReturnStatement is the node for return statement
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral for ReturnStatement
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
