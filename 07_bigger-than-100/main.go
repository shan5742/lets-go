/*
Problem #7: Bigger than 100
Author: Asam Shan
Date Created: 21/05/2019
Comments: adding the '\n' to the end to eliminate the % being printed out in zsh
*/

/*
***************** Bigger than 100 *****************
Write a program that reads a number from the keyboard via the scanf function, and then prints the message The
number is bigger than 100 if it is. If the number is 100 or less than 100 then you should print out the message: The
number is not bigger than 100.
*/

package main

import (
	"fmt"
)

func main() {

	// First job is to create a variable to store the user input
	var input int32

	// Print out instructions for the user, store their input in our variable, we will use the %d specifier again to specify an int
	fmt.Println("Hey, enter a number...")
	fmt.Scanf("%d", &input)

	// Simple if statement to check whether the input is equal to, smaller than or greater than 100
	if input > 100 {
		fmt.Printf("Your number %d is bigger than 100\n", input)
	} else if input <= 100 {
		fmt.Printf("Your number %d is equal to or less than 100\n", input)
	}

}
