// brainfuck
// https://github.com/raklaptudirm/brainfuck
// Copyright (c) 2021 Rak Laptudirm.
// Licensed under the MIT license.

// Package errors provides functions to detect errors.
//
// The following methods are used to detect, print and
// exit when an error occurs.
//
package errors

import (
	"fmt"
	"os"
)

// StrictCheck checks in an error has occurred in
// the argument, and exits if it has.
func StrictCheck(e error) {
	if e != nil {
		fmt.Print(e)
		os.Exit(0)
	}
}

// Check checks if an error has occurred or not,
// and returns true if it has not.
func Check(e error) bool {
	if e != nil {
		fmt.Println(e)
		return false
	}

	return true
}
