package main

import (
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/raklaptudirm/brainfuck/parser"
	. "github.com/raklaptudirm/brainfuck/types"
	. "github.com/raklaptudirm/brainfuck/vm"
)

const help = "Usage: brainfuck <command> [flags]\n\nThe commands are:\n  run      run a  brainfuck file\n  test     test a brainfuck file\n  repl     start the brainfuck repl\n\nUse \"brainfuck help <command>\" for more information about a help."

var VIRTUAL_MACHINE VM = VM{DataPointer: 0, Memory: [TAPE_LENGTH]byte{}}

func strictCheck(e error) {
	if e != nil {
		fmt.Print(e)
		os.Exit(0)
	}
}

func check(e error) bool {
	if e != nil {
		fmt.Println(e)
		return false
	}

	return true
}

func main() {
	brainfuck := VIRTUAL_MACHINE
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(help)
	} else if args[0] == "run" {
		file, fileError := ioutil.ReadFile(args[1])
		strictCheck(fileError)

		instructions, parseError, indexes := Parse(string(file))
		strictCheck(parseError)

		brainfuck.Run(os.Stdout, instructions, indexes)
	} else if args[0] == "repl" {
		fmt.Println("Welcome to brainfuck v1.0.0.")
		for {
			fmt.Print("> ")
			var input string
			fmt.Scanln(&input)

			if input == "exit" {
				break
			} else {
				instructions, parseError, indexes := Parse(input)

				if check(parseError) {
					brainfuck.Run(os.Stdout, instructions, indexes)
				}

				length := len(instructions)
				for i := 0; i < length; i++ {
					if instructions[i] == OUTPUT {
						fmt.Print("\n")
						break
					}
				}
			}
		}
	} else if args[0] == "test" {
		file, fileError := ioutil.ReadFile(args[1])
		strictCheck(fileError)

		_, parseError, _ := Parse(string(file))
		strictCheck(parseError)

		fmt.Printf("%v: No errors found.", args[1])
	} else if args[0] == "help" {
		switch args[1] {
		case "run":
			fmt.Println("Usage: brinfuck run <file> [flags]\n\nRuns a brainfuck file.\nIt tries to parse the file and,\nif successful, runs it.\n\nFlags:\n  to be added")
		case "test":
			fmt.Println("Usage: brainfuck test <file> [flags]\n\nTests if a file is parsable.\nIt tries to parse the file, and then exits.\n\nFlags:\n  to be added")
		case "repl":
			fmt.Println("Usage: brainfuck repl [flags]\n\nStarts the brainfuck repl.\n\nFlags:\n  to be added")
		default:
			fmt.Printf("error: unknown command %v.", args[1])
		}
	} else {
		fmt.Printf("error: unknown command %v.", args[0])
	}
}
