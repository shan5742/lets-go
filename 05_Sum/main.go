/*
Problem #5: The sum of two values
Author: Asam Shan
Date Created: 21/05/2019
*/

/*
***************** The sum of two values *****************
Write a program that reads two integer values from the keyboard via the scanf function, then adds them together,
stores the result into a variable called sum, and then prints out the value of the variable sum.
*/

package main

import "fmt"

/* This is somewhat similar to problem 4, we must first create 2 variables, num1 and num2, which the user will input, we will again utilise the Scanf function for this. We then need to add those numbers togetehr and print out the result */

var num1, num2, sum int // ccreate our variables and set them as int

func main() {
	// Firstly we give instructions for the user input.
	fmt.Printf("Hey, pick a number...")

	// Now we grab that number and assign it to our num1 variable earlier varaible, using the %d specifier for an int"
	fmt.Scanf("%d", &num1)

	// Now we give follow up instructions for the user
	fmt.Print("Give me another number and I'll add them together for you...")

	// Now we grab that number and assign it to our num2 variable earlier varaible"
	fmt.Scanf("%d", &num2)

	// Now we add the simple logic of the sum
	sum = num1 + num2

	// And finally we print out the the sum, again using the %d to specify an int
	fmt.Printf("Thanks, the sum of your numbers is:%d\n", sum)

}
