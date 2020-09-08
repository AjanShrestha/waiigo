package ast

import (
	"bytes"
	"monkey/token"
)

// Node is the base Node interface
type Node interface {
	TokenLiteral() string
	String() string
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
	}
	return ""
}

// String returns Program Node
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Identifier node
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral for Identifiers
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String returns the Identifier Node
func (i *Identifier) String() string {
	return i.Value
}

// LetStatement is the node for let statement
type LetStatement struct {
	Token token.Token // the token.Let Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral for LetStatement
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// String returns the Let Node
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement is the node for return statement
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral for ReturnStatement
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// String returns the Return Node
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is the node for expressions
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral for ExpressionStatement
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String returns the Expression node
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// IntegerLiteral defines integers
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral returns the node value
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String returns the node repr
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
