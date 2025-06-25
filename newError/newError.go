/*
In Go (Golang), nil is a predeclared identifier representing zero value for 
pointers, interfaces, maps, slices, channels, and function types. 
It essentially means "nothing" or "no value."
*/

package main

import (
	"errors"
	"fmt"
)

func returnError(a,b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!!!")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1,2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally")
	} else {
		fmt.Println(err)
	}

	if err.Error() == "Error in returnError( function)" {
	fmt.Println("!!")
	}
}


/*
The code prints the error message on the screen and exits using os.Exit().
if err != nil {
 log.Println(err)
 os.Exit(10)
}
*/

/*
The code that is used when something really
bad has happened and you want to terminate the program:
if err != nil {
 panic(err)
 os.Exit(10)
}

*/

/*
The panic function is a built-in Go function that stops the execution of a program and
starts panicking! If you find yourself using panic too often, you might want to reconsider
your Go implementation.
*/

/*
Go alsooffers the recover function, 
which might be able to save you when you're in some bad situations.
*/