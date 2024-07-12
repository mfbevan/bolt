package ast

import "bolt/token"

// Node is the interface that all nodes in the AST implement.
//   - TokenLiteral: returns the literal value of the token
type Node interface {
	TokenLiteral() string
}

// Statement is the interface that all statement nodes in the AST implement.
//   - statementNode: a marker for statement nodes
type Statement interface {
	Node
	statementNode()
}

// Expression is the interface that all expression nodes in the AST implement.
//   - expressionNode: a marker for expression nodes
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of every AST that the parser produces.
//   - Statements: a slice of statements in the program
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal value of the first statement in the program.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement represents a let statement in the AST.
//   - Token: the token.LET token
//   - Name: the identifier of the let statement
//   - Value: the expression that the let statement is bound to
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier represents an identifier in the AST.
//   - Token: the token.IDENT token
//   - Value: the value of the identifier
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ReturnStatement represents a return statement in the AST.
//   - Token: the token.RETURN token
//   - ReturnValue: the expression that the return statement is bound to
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
