/*
To take input there are two ways to do this:
1. Using fmt.Scanln() method like:
var input string
fmt.Scanln(&input)
It scans text read from standard input, storing successive space-separated values into the successive arguments.

2. Using bufio.NewReader() method like:
reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n')

This creates a new buffered reader that takes input from the standard input, os.Stdin.

Underscore ( "_" ) is used widely in comman, okay syntax with := operator as a placeholder.
The underscore _ is a blank identifier, used when syntax requires a variable name but program logic does not.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || err err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)

	/*
		We can also take inputs using fmt.Scanln() instead of using bufio.NewReader()
		var input string
		fmt.Scanln(&input)
	*/

}
