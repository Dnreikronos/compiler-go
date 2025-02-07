func Generate(expr parser.Expr) string {
	output := ".global_start\n"
	output += "_start:\n"
	output += traverse(expr)
	output += "		mov rax, 60\n"
	output += "		mov rdi, 0\n"
	output += "		syscall\n"
	return output
}
