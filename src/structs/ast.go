package structs

// replace below w/ these structs:
/*
type ASTNode struct {
    Type     string
    Value    string
    Children []*ASTNode
}

type Program struct {
    Children []*ASTNode
}

type ExpressionStatement struct {
    Expression *ASTNode
}

type CallExpression struct {
    Caller    *ASTNode
    Arguments []*ASTNode
}

type MemberExpression struct {
    Object  *ASTNode
    Property *ASTNode
}

type StringLiteral struct {
    Value string
}
*/

type ASTNode interface {
    doNothing()
}

type Program struct {
    Type string
    Body []ASTNode
}

func NewProgram() Program {
    return Program{Type: "Program"}
}

func (p Program) doNothing() {
}

type VariableDeclaration struct {
    Type string
    Kind string
    Declarations []ASTNode
}

func NewVariableDeclaration() VariableDeclaration {
    return VariableDeclaration{Type: "VariableDeclaration"}
}

func (vdn VariableDeclaration) doNothing() {
}

type ExpressionStatement struct {
    Type string
    Expression ASTNode
}

func NewExpressionStatement() ExpressionStatement {
    return ExpressionStatement{Type: "ExpressionStatement"}
}

func (exs ExpressionStatement) doNothing() {
}

type CallExpression struct {
    Type string
    Callee ASTNode
    Arguments ASTNode
}

func NewCallExpression() CallExpression {
    return CallExpression{Type: "CallExpression"}
}

func (cex CallExpression) doNothing() {
}

type Identifier struct {
    Type string
    Identifier string
}

func NewIdentifier(identifier string) Identifier {
    return Identifier{Type: "Identifier", Identifier: identifier}
}

func (idf Identifier) doNothing() {
}

type MemberExpression struct {
    Type string
    Object ASTNode
    Property ASTNode
}

type VariableDeclarator struct {
    Type string
    Id ASTNode
    Init ASTNode
}

func newVariableDeclarator() VariableDeclarator {
    return VariableDeclarator{Type: "Variable Declarator"}
}

func (vdr VariableDeclarator) doNothing() {
}

type NumericLiteral struct {
    Type string
    Value float64
}

func NewNumericLiteral() NumericLiteral {
    return NumericLiteral{Type: "NumericLiteral"}
}

func (nl NumericLiteral) doNothing() {
}

type StringLiteral struct {
    Type string
    Value string
}

func NewStringLiteral(stringLiteral string) StringLiteral {
    return StringLiteral{Type: "StringLiteral", Value: stringLiteral}
}

func (sl StringLiteral) doNothing() {
}

type Punctuation struct {
    Type string
    Value string
}

func NewPunctuation(literal string) Punctuation {
    return Punctuation{Type: "punctuation", Value: literal}
}

func (p Punctuation) doNothing() {
}
