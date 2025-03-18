package tokenizer

import (
    "testing"
)

func TestTokenize(t *testing.T) {
    lines := []string{`console.log("foo")`}
    result := Tokenize(lines)
    if (result[0] != "console" || result[1] != "log" || result[2] != "foo") {
        t.Error("operation failed")
    }
}
