package astgen

import (
    "testing"

    "de.asiegwarth/jscomp/src/structs"
)

func TestCreateASTNodesFromTokens(t *testing.T) {
    astNodes, err := createASTNodes("../../data/consolelog.js")
    if err != nil {
        t.Error("no AST nodes created")
    }
    _, ok := astNodes.Body[0].(structs.Identifier)
    if ok == false {
        t.Error("wrong or no AST nodes created")
    }
    var value interface{}
    value, ok = astNodes.Body[1].(structs.Punctuation)
    if ok == false {
        t.Errorf("wrong or no AST nodes created: %v\n", value)
    }
}
