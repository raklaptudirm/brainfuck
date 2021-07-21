package main

import (
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/raklaptudirm/brainfuck/parser"
	. "github.com/raklaptudirm/brainfuck/types"
	. "github.com/raklaptudirm/brainfuck/vm"
)

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
	} else {
		file, fileError := ioutil.ReadFile(args[0])
		strictCheck(fileError)

		instructions, parseError, indexes := Parse(string(file))
		strictCheck(parseError)

		brainfuck.Run(os.Stdout, instructions, indexes)
	}
}
