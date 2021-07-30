package errors

import (
	"fmt"
	"os"
)

func StrictCheck(e error) {
	if e != nil {
		fmt.Print(e)
		os.Exit(0)
	}
}

func Check(e error) bool {
	if e != nil {
		fmt.Println(e)
		return false
	}

	return true
}
