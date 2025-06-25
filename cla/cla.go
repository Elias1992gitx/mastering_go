/*
The program will find the minimum and the
maximum of its command-line arguments
*/

/* To obtain the command-line arguments requires
the use of the os package */

/* 'strconv' order to be able to convert a command-line argument, which is given as a string, into an
arithmetical data type. */

/*
the error ignored returned by the strconv.ParseFloat() function
using the following statement:
n, _ := strconv.ParseFloat(arguments[i], 64)


The preceding statement tells Go that you only want to get the first value returned by
strconv.ParseFloat() and that you are not interested in the second value, which in this
case is an error variable by assigning it to the underscore character. The underscore
character, which is called a blank identifier, is the Go way of discarding a value. If a Go
function returns multiple values, you can use the blank identifier multiple times.
*/


package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	min, _:= strconv.ParseFloat(arguments[1], 64)
	max, _:= strconv.ParseFloat(arguments[1], 64)
	
	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("min:", min)
	fmt.Println("max:", max)
}


/*
WARNING: Ignoring all or some of the return values of a Go function,
especially the error values, is a very dangerous practice that should not be
used in production code!

*/