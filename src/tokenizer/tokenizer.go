package tokenizer

import (
    "fmt"
)

func Tokenize(code []string) []string {
    var tokens []string

    for _, line := range code {
        fmt.Println(line)
    }

    return tokens
}
