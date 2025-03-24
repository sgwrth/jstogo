package astgen

import (
    "errors"

    "de.asiegwarth/jscomp/src/io"
    "de.asiegwarth/jscomp/src/structs"
    "de.asiegwarth/jscomp/src/tokenizer"
)

func createASTNodes(codepath string) (*structs.Program, error) {
    code, err := io.ReadJSFile(codepath)
    if err != nil {
        return nil, errors.New("error reading code file")
    }
    tokens := tokenizer.Tokenize(code)
    astNodes := structs.Program{}
    for _, token := range tokens {
        tokenType := tokenizer.DetermineTokenType(token)
        switch tokenType.(type) {
        case structs.IdentifierToken:
            astNodes.Body = append(astNodes.Body, structs.NewIdentifier(token))
        case structs.StringLiteralToken:
            astNodes.Body = append(astNodes.Body, structs.NewStringLiteral(token))
        }
    }
    return &astNodes, nil
}
