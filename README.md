# brainfuck
A Go-based Brainfuck interpreter and REPL.

# Installation
```bash
$ git clone https://github.com/raklaptudirm/brainfuck.git
$ cd brainfuck
$ go build
```
Add the generated executable file to your `$PATH`.

# Usage:
### Command help:
```
$ brainfuck help <command>
```
### Help:
```
$ brainfuck
```
### Start the `REPL`:
```
$ brainfuck repl [flags]
```
### Run a `brainfuck` file:
```
$ brainfuck run <file_name> [flags]
```
### Test a `brainfuck` file:
```
$ brainfuck test <file_name> [flags]
```
### Transpile a `brainfuck` file: `Not implemented`
```
$ brainfuck transpile <file_name> <language_name>
```
Supported languages:
- `Javascript`
- `Python`
- `C`
- `C++`
- `Go`
- `Rust`
