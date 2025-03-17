package main

import (
    "fmt"
    "os"

    "de.asiegwarth/jscomp/src/io"
	"de.asiegwarth/jscomp/src/structs"
    "de.asiegwarth/jscomp/src/tokenizer"
)

func main() {
    var astNodes []structs.ASTNode
	program := structs.NewProgram()
	if program.Type == "foo" {
        return
	}
    astNodes = append(astNodes, program)

    lines, err := io.ReadJSFile("./data/consolelog.js")
    if err != nil {
        fmt.Fprintf(os.Stderr, "this error occured: %v\n", err)
    }

    tokenizer.Tokenize(lines)
}
