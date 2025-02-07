func Generate(expr parser.Expr) string {
	output := ".global_start\n"
	output += "_start:\n"
	output += traverse(expr)
	output += "		mov rax, 60\n"
	output += "		mov rdi, 0\n"
	output += "		syscall\n"
	return output
}
func traverse(expr parser.Expr) string {
	switch e := expr.(type) {
	case parser.NumberExpr:
		return fmt.Sprintf("    mov rax, %s\n", e.Value)
	case parser.BinaryExpr:
		left := traverse(e.Left)
		right := traverse(e.Right)
		code := left +
			"    push rax\n" +
			right +
			"    pop rbx\n"
		switch e.Op {
		case "+":
			code += "    add rax, rbx\n"
		case "-":
			code += "    sub rax, rbx\n"
		case "*":
			code += "    imul rax, rbx\n"
		case "/":
			code += "    cqo\n"
			code += "    idiv rbx\n"
		}
		return code
	default:
		panic("Unsupported expression")
	}
}
