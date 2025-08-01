package main

import (
	"go/parser"
	"go/token"
)

func syntaxValid(src string) bool {
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "", src, parser.AllErrors)
	return err == nil
}
