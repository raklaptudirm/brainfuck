/*
* brainfuck
* https://github.com/raklaptudirm/brainfuck
*
* Copyright (c) 2021 Rak Laptudirm
* Licensed under the MIT license.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/raklaptudirm/brainfuck/help"

	. "github.com/raklaptudirm/brainfuck/vm"
	. "github.com/raklaptudirm/brainfuck/errors"
	. "github.com/raklaptudirm/brainfuck/parser"
)

func main() {
	brainfuck := VMBase
	args := os.Args[1:]

	assert := func (length int) {
  	if len(args) != length + 1 {
  		fmt.Printf("Expected %v arg(s), received %v.", length, len(args) - 1)
  		os.Exit(0)
  	}
  }

	if len(args) == 0 {
		fmt.Println(help.Default)
	} else {
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
			StrictCheck(fileError)

			_, parseError, _, _ := Parse(string(file))
			StrictCheck(parseError)

			fmt.Printf("%v: No errors found.", args[1])
		default:
			fmt.Printf("error: unknown command %v.", args[0])
		}
	}
}
