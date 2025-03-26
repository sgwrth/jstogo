package tokenizer

import (
    "testing"

    "de.asiegwarth/jscomp/src/structs"
)

func TestTokenize(t *testing.T) {
    lines := []string{`console.log("foo");`}
    result := Tokenize(lines)
    if (result[0] != "console" ||
            result[1] != "." ||
            result[2] != "log" ||
            result[3] != "(" ||
            result[4] != `"` ||
            result[5] != "foo" ||
            result[6] != `"` ||
            result[7] != ")" ||
            result[8] != ";") {
        t.Error("operation failed")
    }
}

func TestTokenType(t *testing.T) {
    token := DetermineTokenType("let")
    if token == nil {
        t.Error("no token type available")
    }
    if token.(structs.KeywordToken).Type != "keyword" {
        t.Error("wrong token type")
    }
    token = DetermineTokenType("console")
    if token == nil {
        t.Error("no token type available")
    }
    if token.(structs.IdentifierToken).Type != "identifier" {
        t.Error("wrong token type")
    }
    token = DetermineTokenType(`"`)
    if token.(structs.PunctuationToken).Type != "punctuation" {
        t.Error("wrong token type")
    }
}
