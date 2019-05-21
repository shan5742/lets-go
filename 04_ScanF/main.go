/*
Problem #4: The Scan F function
Author: Asam Shan
Date Created: 21/05/2019
*/

/*
***************** The Scan F function *****************
Write a program that reads in an integer value from the keyboard via the scanf function and then prints it back
onto the screen. You should know that scanf seeks an address expression. For example, &n is an address expression.
You should also know that scanf returns a value. It is either the number of successful conversions or an error code
that is typically EOF.
*/

package main

import "fmt"

var name string // creating a variable for the user input, in the case the users name

func main() {
	// Firstly we give instructions for the user input.
	fmt.Printf("Hey, what's your name?")

	// Now we use the Scanf method to record teh users input, we are using the %s to take in a string and attach it to our earlier varaible with "&name"
	fmt.Scanf("%s", &name)

	// And finally we print out the user input, using the PrintF again, as you can see we have used the %s again to specify a string and also have included the \n for a new line
	fmt.Printf("Cool, your name is:%s\n", name)

}
