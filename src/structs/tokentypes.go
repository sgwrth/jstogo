package structs

type Token struct {
    Type string
    Literal string
}

func NewToken(tokenType string, literal string) Token {
    return Token{Type: tokenType, Literal: literal}
}

type KeywordToken struct {
    Token
}

func NewKeywordToken(literal string) KeywordToken {
    return KeywordToken{Token: NewToken("keyword", literal)}
}

type IdentifierToken struct {
    Token
    Name string
}

func NewIdentifierToken(literal string) IdentifierToken {
    return IdentifierToken{Token: NewToken("identifier", literal), Name: literal}
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

func NewPunctuationToken(literal string) PunctuationToken {
    return PunctuationToken{Token: NewToken("punctuation", literal)}
}
