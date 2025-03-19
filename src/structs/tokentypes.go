package structs

type Token struct {
    Type string
    Literal string
}

type KeywordToken struct {
    Token
}

type IdentifierToken struct {
    Token
    Name string
}

type StringLiteralToken struct {
    Token
    Value string
}

type NumericLiteralToken struct {
    Token
    Value float64
}

type OperatorToken struct {
    Token
}

type PunctuationToken struct {
    Token
}
