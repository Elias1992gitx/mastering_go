
package main

import (
	"io"
	"os"
)
func main() {
	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "Please give me one arguments!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is Standared output \n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
