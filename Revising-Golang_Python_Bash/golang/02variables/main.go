/*
Key Concepts covered in this program
In Go, you can declare variables by using the var keyword, followed by the variable name and the variable type. For example=>  var smallVal uint8 = 255
Go also supports shorthand variable declaration with the walrus := operator. For example=>  numberOfUser := 300000

While printing variables in golang and using fmt.Printf() function, these are the placeholders for various types:
%d  =>  Integer  (int)
%s  =>  String   (string)
%t  =>  Bool     (bool)
%f  =>  Flaot64  (float64)

Access Specifiers: Go doesn't have explicit keywords for declaring variables public or private. Instead, Go uses a concept called exported and unexported identifiers.
An identifier is exported if it begins with an uppercase letter.
An identifier is unexported if it begins with a lowercase letter.

*/

package main

import "fmt"

const LoginToken string = "ghabbhhjd" // Public variable

func main() {
	var firstName string = "yuvraj"
	lastName := "gupta"
	fmt.Printf("First Name: %s\nLast Name: %s\n", firstName, lastName)
	fmt.Printf("Variable is of type: %T \n", firstName)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var intVar int
	var stringVar string
	var boolVar bool
	var floatVar float64

	fmt.Printf("Default value of interger %d \nDefault value of string: %s \nDefault value of bool: %t \nDefault value of float: %f\n", intVar, stringVar, boolVar, floatVar)

	// implicit type

	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
