/*The os package is used
for reading the command-line arguments of the program and 
for accessing os.Stdout.*/






package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}


/*the io.WriteString() function works in the same way as the fmt.Print()
function; however, it takes only two parameters. The first parameter is the file to which you
want to write, which in this case is os.Stdout, and the second parameter is a string
variable*/


/*
Getting user input
There are three main ways to acquire user input:
1. By reading the command-line arguments of a program
2. By asking the user for input
3. By reading external files
*/