package main

import (
	"de.asiegwarth/jscomp/src/structs"
)

func main() {
    var astNodes []structs.ASTNode
	program := structs.NewProgram()
	if program.Type == "foo" {
        return
	}
    astNodes = append(astNodes, program)
}
