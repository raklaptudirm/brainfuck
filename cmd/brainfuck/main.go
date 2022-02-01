// brainfuck
// https://github.com/raklaptudirm/brainfuck
// Copyright (c) 2021 Rak Laptudirm.
// Licensed under the MIT license.

// Brainfuck is a cli compiler and repl
// for the brainfuck language.
//
// Usage:
//
//      brainfuck <command> [flags]
//
//      The commands are:
//        run      run a  brainfuck file
//        test     test a brainfuck file
//        repl     start the brainfuck repl
//
// Example:
//
// Run a simple hello world program from
// the examples directory.
//
//      $ brainfuck run examples/hello_world.bf
//      Hello World!
//
// Test a brainfuck file for parse errors.
//
//      $ brainfuck test examples/hello_world.bf
//      examples/hello_world.bf: No errors found.
//
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/raklaptudirm/brainfuck/pkg/help"

	"github.com/raklaptudirm/brainfuck/pkg/errors"
	"github.com/raklaptudirm/brainfuck/pkg/parser"
	"github.com/raklaptudirm/brainfuck/pkg/vm"
)

func main() {
	brainfuck := vm.Default
	args := os.Args[1:]

	// assert function to check wether the expected number
	// of arguments was received or not.
	assert := func(length int) {
		if len(args) != length+1 {
			fmt.Printf("Expected %v arg(s), received %v.", length, len(args)-1)
			os.Exit(0)
		}
	}

	if len(args) == 0 {
		fmt.Println(help.Default)
	} else {
		// args[0] contains the command name.
		switch args[0] {
		case "run":
			assert(1)
			brainfuck.RunFile(args[1], os.Stdout)
		case "help":
			assert(1)
			help.Get(args[1])
		case "repl":
			fmt.Println("The new REPL is a work in progress.")
		case "test":
			assert(1)
			file, fileError := ioutil.ReadFile(args[1])
			errors.StrictCheck(fileError)

			_, parseError, _ := parser.Parse(string(file))
			errors.StrictCheck(parseError)

			fmt.Printf("%v: No errors found.", args[1])
		default:
			fmt.Printf("error: unknown command %v.", args[0])
		}
	}
}
