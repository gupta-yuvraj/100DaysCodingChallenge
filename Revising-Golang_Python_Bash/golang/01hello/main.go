/*
Let's break down this code:

`package main` : This line declares that this code is in the main package. The main package is special in Go because it defines the entry point for the executable program.

`import "fmt"` : This line imports the fmt package, which is a standard library package for formatting input and output. It is commonly used to print values to the screen.

`func main()` { : This line declares the main function, which is the entry point for the program. When the program runs, the code inside this function will be executed.

`fmt.Println("Hello World")` : This line is a statement inside the main function. It calls the Println function from the fmt package to print the string "Hello World" followed by a newline to the standard output (usually the terminal screen).

`}` : This line closes the main function.
*/


package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

