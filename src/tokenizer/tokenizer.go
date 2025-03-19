package tokenizer

import (
    "fmt"
    "regexp"

    "de.asiegwarth/jscomp/src/structs"
)

const specialCharRegExp = `(\s|\(|\)|\;|\.)` // ' ', '(', ')', ';', '.'

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

func DetermineTokenType(token string) structs.Token {
    return structs.Token{}
}
