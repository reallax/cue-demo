package main

import (
	"fmt"

	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/parser"
)

func main() {
	// Parse the file containing this very example
	// but stop after processing the imports.
	f, err := parser.ParseFile("test.cue", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the imports from the file's AST.
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}
	
	fmt.Println(f.Filename)
	fmt.Println(ast.Name(f.Decls[2]))
	fmt.Println(f.Comments())
	fmt.Println(f.Unresolved)
}