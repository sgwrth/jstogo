package tokenizer

import (
    "fmt"
    "regexp"

    "de.asiegwarth/jscomp/src/dict"
    "de.asiegwarth/jscomp/src/structs"
)

const specialCharRegExp = `(\(|\)|;|\.|")` // punctuation: () . ;

func Tokenize(code []string) []string {
    var allTokens []string

    for _, line := range code {
        fmt.Println(line)
        var token string
        for _, char := range line {
            if char == ' ' {
                continue
            }
            match, _ := regexp.MatchString(specialCharRegExp, string(char))
            if match {
                if token != "" {
                    allTokens = append(allTokens, token)
                }
                allTokens = append(allTokens, string(char))
                token = ""
                continue
            } else {
                token += string(char)
            }
        }
    }

    for _, token := range allTokens {
        fmt.Println(token)
    }        

    return allTokens
}

func DetermineTokenType(token string) interface{} {
    for _, keyword := range dict.Keywords {
        if keyword == token {
            // return structs.KeywordToken{Token: structs.Token{Type: "keyword", Literal: token}}
            return structs.NewKeywordToken(token)
        }
    }

    for _, punctuation := range dict.Punctuation {
        if punctuation == token {
            // return structs.PunctuationToken{Token: structs.Token{Type: "punctuation", Literal: token}}
            return structs.NewPunctuationToken(token)
        }
    }

    // return structs.IdentifierToken{Token: structs.Token{Type: "identifier", Literal: token}}
    return structs.NewIdentifierToken(token)
}
