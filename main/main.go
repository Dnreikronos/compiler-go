package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Dnreikronos/compiler-go/codegen"
	"github.com/Dnreikronos/compiler-go/lexer"
	"github.com/Dnreikronos/compiler-go/parser"
)
func main() {
	input := "1 + 2 * 3"

	tokens := lexer.Lex(input)

	p := parser.NewParser(tokens)
	ast := p.Parse()

	asm := codegen.Generate(ast)

	os.WriteFile("output.s", []byte(asm), 0644)

	cmd := exec.Command("nasm", "-felf64", "output.s")
	cmd.Run()
	cmd = exec.Command("ld", "-o", "output", "output.o")
	cmd.Run()

	fmt.Println("Compilation successful. Run './output'")
}
