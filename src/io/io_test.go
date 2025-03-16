package io

import (
    "testing"
)

func TestReadJSFile(t *testing.T) {
    filepath := "../../data/consolelog.js"
    actualContent := []string{`console.log("foo");`}

    readContent, err := ReadJSFile(filepath)
    if err != nil {
        t.Error(`Error reading the file`)
    }
    if readContent[0] != actualContent[0] {
        t.Error(`Error reading file content`)
    }
}
