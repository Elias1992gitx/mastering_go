package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File       //Declare a file pointer

	f = os.Stdin         //set input source to keyboard

	defer f.Close()      //defer means "run this later" - at the end of the function

	scanner := bufio.NewScanner(f)         //Create scanner for reading input lines
	for scanner.Scan() {                   //Check for and read next input line
		fmt.Println(">", scanner.Text ())  //Get the content of that input line
	}
}










/*The official name for := is the short assignment statement. The short assignment
statement can be used in place of a var declaration with an implicit type*/

// package main

// import (
// 	"fmt"
// 	// "log"
// )

// func main() {
// 	x, y := 2, 5
// 	l, y := 2, 7

// 	fmt.Println("The Vaue of 'X'= ", x)
// 	fmt.Println("The Vaue of 'y'= ", y)
// 	fmt.Println("The Vaue of 'l'= ", l)
// }
