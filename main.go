package main

import (
	"fmt"

	"github.com/QuantumWizd/go-complier/lexer"
)

func main() {

	fmt.Println("welcome to go compiler ")

	input := `variable x = 10
constant y = 20

function add(x, y) ~ {
    return x + y
}

function main() {
    result :: add(10, 20)
    
    if result == 30 {
        print result
    }

    for {
        
    }
}`

	l := lexer.NewLexer(input)

	tokens := l.Tokenize()
	for _, tok := range tokens {
		fmt.Println(tok)
	}

}
