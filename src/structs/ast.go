package structs

type ASTNode interface {
    doNothing()
}

type Program struct {
    Type string
    Body []ASTNode
}

func NewProgram() Program {
    return Program{
        Type: "Program",
    }
}

func (p Program) doNothing() {
}

type VariableDeclaration struct {
    Type string
    Kind string
    Declarations []ASTNode
}

func NewVariableDeclaration() VariableDeclaration {
    return VariableDeclaration{
        Type: "VariableDeclaration",
    }
}

func (vdn VariableDeclaration) doNothing() {
}

type ExpressionStatement struct {
    Type string
    Expression ASTNode
}

func NewExpressionStatement() ExpressionStatement {
    return ExpressionStatement{
        Type: "ExpressionStatement",
    }
}

func (exs ExpressionStatement) doNothing() {
}

type CallExpression struct {
    Type string
    Callee ASTNode
    Arguments ASTNode
}

func NewCallExpression() CallExpression {
    return CallExpression{
        Type: "CallExpression"
    }
}

func (cex CallExpression) doNothing() {
}

type Identifier struct {
    Type string
    Identifier string
}

func NewIdentifier() Identifier {
    return Identifier{
        Type: "Identifier"
    }
}

func (idf Identifier) doNothing {
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
    return VariableDeclarator{
        Type: "Variable Declarator"
    }
}

func (vdr VariableDeclarator) doNothing() {
}

type NumericLiteral struct {
    Type string
    Value float64
}

func NewNumericLiteral() NumericLiteral {
    return NumericLiteral{
        Type: "NumericLiteral"
    }
}

func (nl NumericLiteral) doNothing() {
}
