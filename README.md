# Golang Arithmetic Expression Compiler

## Overview
This project is a simple compiler for arithmetic expressions written in Golang. It includes a lexer, parser, and code generator that produces x86-64 assembly code, which can be assembled and linked to create an executable.

## Features
- Tokenizes arithmetic expressions (Lexer)
- Parses expressions into an Abstract Syntax Tree (AST) (Parser)
- Generates x86-64 assembly code (Code Generator)
- Supports basic arithmetic operations: `+`, `-`, `*`, and `/`
- Uses NASM for assembly and `ld` for linking

## Project Structure

compiler-go/
├── lexer/        # Lexical analysis (tokenizer) <br>
├── parser/       # Parsing logic <br>
├── codegen/      # Code generation <br>
├── main.go       # Entry point of the compiler <br>
├── output.s      # Generated assembly file <br>
├── output        # Compiled binary <br>


## Installation
Ensure you have Golang, NASM, and a linker installed:

sh
sudo apt install nasm
sudo apt install gcc  # Includes ld (the linker)


Clone the repository:
sh
git clone https://github.com/Dnreikronos/compiler-go.git
cd compiler-go


## Usage
Run the compiler with:
sh
go run main.go

This will:
1. Tokenize and parse the hardcoded expression (currently `1 + 2 * 3`).
2. Generate assembly code in `output.s`.
3. Assemble the code into `output.o` using NASM.
4. Link the object file into an executable `output`.

Run the generated binary:
sh
./output


## Example
Given the input:

1 + 2 * 3

The generated assembly (`output.s`) will be similar to:

.global _start
_start:
    mov rax, 2
    push rax
    mov rax, 3
    pop rbx
    imul rax, rbx
    push rax
    mov rax, 1
    pop rbx
    add rax, rbx
    mov rax, 60
    mov rdi, 0
    syscall


## Contributing
Feel free to fork the repository and submit pull requests for improvements!


