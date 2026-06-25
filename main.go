package main

import (
	"fmt"

	"github.com/QuantumWizd/go-complier/lexer"
)

func main() {

	fmt.Println("welcome to go compiler ")

	input := `let x = 10
let y = 20
print x`

	l := lexer.NewLexer(input)
	
	tokens := l.Tokenize()
	for _, tok := range tokens {
		fmt.Println(tok)
	}

}
