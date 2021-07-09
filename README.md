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
### Run a `brainfuck` file:
```
$ brainfuck <file_name>
```
For example:
```
$ brainfuck examples/mandelbrot.bf
```

### Start the `REPL`:
```
$ brainfuck
```
# Later API (To be implemented):
### Command help:
```
$ brainfuck
```
### Start the `REPL`:
```
$ brainfuck repl
```
### Run a `brainfuck` file:
```
$ brainfuck run <file_name>
```
### Transpile a `brainfuck` file:
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
