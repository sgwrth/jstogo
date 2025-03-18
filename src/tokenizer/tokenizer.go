package tokenizer

import (
    "fmt"
    "regexp"
)

const specialCharRegExp = `(\s|\(|\)|\;|\.|\")` // ' ', '(', ')', ';', '.', '"'

func Tokenize(code []string) []string {
    var allTokens []string

    for _, line := range code {
        var token string
        for _, ch := range line {
            match, _ := regexp.MatchString(specialCharRegExp, string(ch))
            if match {
                if token != "" {
                    allTokens = append(allTokens, token)
                }
                token = ""
                continue
            } else {
                token += string(ch)
            }
        }
    }

    for _, token := range allTokens {
        fmt.Println(token)
    }        

    return allTokens
}
