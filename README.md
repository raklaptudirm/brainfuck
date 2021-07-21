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
## Help:
Get help about the cli app.
```
$ brainfuck
```
## Command help:
Get help about a specific command.
```
$ brainfuck help <command>
```
## Start the `REPL`:
Start the `brainfuck` repl provided by the cli.
```
$ brainfuck repl [flags]
```
## Run a `brainfuck` file:
Parse and execute a `brainfuck` file.
```
$ brainfuck run <file_name> [flags]
```
## Test a `brainfuck` file:
Parse a `brainfuck` file and check for errors.
```
$ brainfuck test <file_name> [flags]
```
## Transpile a `brainfuck` file: `Not implemented`
```
$ brainfuck transpile <file_name> <language_name>
```
### Supported languages:
- `Javascript`
- `Python`
- `C`
- `C++`
- `Go`
- `Rust`
