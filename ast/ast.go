package ast

import (
	"bytes"
	"monkey/token"
)

// Types
type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program
type Program struct {
	Statements []Statement
}

func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	}

	return ""
}

type ExpressionStatement struct {
	Statement
	Token      token.Token // token.???
	Expression Expression
}

func (ls *ExpressionStatement) statementNode()       {}
func (ls *ExpressionStatement) TokenLiteral() string { return ls.Token.Literal }

type ReturnStatement struct {
	Statement
	Token token.Token // token.RETURN
	ReturnValue Expression
}

func (ls *ReturnStatement) statementNode()       {}
func (ls *ReturnStatement) TokenLiteral() string { return ls.Token.Literal }

type LetStatement struct {
	Statement
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (i *Identifier) String() string {
  return i.Value;
}

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

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

  out.WriteString(rs.TokenLiteral() + " ")

  if rs.ReturnValue != nil {
    out.WriteString(rs.ReturnValue.String())
  }

  out.WriteString(";")

	return out.String()
}

func (ex *ExpressionStatement) String() string {
	if ex.Expression != nil {
		return ex.Expression.String()
	}

	return ""
}
