package help

import (
	"fmt"
)

const Default string = "Usage: brainfuck <command> [flags]\n\nThe commands are:\n  run      run a  brainfuck file\n  test     test a brainfuck file\n  repl     start the brainfuck repl\n\nUse \"brainfuck help <command>\" for more information about a command."

func Get(query string) {
	switch query {
	case "run":
		fmt.Println("Usage: brinfuck run <file> [flags]\n\nRuns a brainfuck file.\nIt tries to parse the file and,\nif successful, runs it.\n\nFlags:\n  to be added")
	case "test":
		fmt.Println("Usage: brainfuck test <file> [flags]\n\nTests if a file is parsable.\nIt tries to parse the file, and then exits.\n\nFlags:\n  to be added")
	case "repl":
		fmt.Println("Usage: brainfuck repl [flags]\n\nStarts the brainfuck repl.\n\nFlags:\n  to be added")
	default:
		fmt.Printf("error: unknown command %v.", query)
	}
}