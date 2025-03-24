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
    if astNodes.Body[0].(structs.Identifier).Type != "Identifier" {
        t.Error("wrong or no AST nodes created")
    }
}
