package tokenizer

import (
    "testing"

    "de.asiegwarth/jscomp/src/structs"
)

func TestTokenize(t *testing.T) {
    lines := []string{`console.log("foo")`}
    result := Tokenize(lines)
    if (result[0] != "console" || result[1] != "log" || result[2] != `"foo"`) {
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
}

