/*
Problem #2: The Value 6
Author: Asam Shan
Date Created: 21/05/2019
comments: Using the ZSH shell omitting the '\n' priunted out an actual % symbol as well as the expected number, the fix was specifying a new line.
*/

/*
***************** The Value 6 *****************
Write a program that prints out the value 6. I want you to use the %d format code.
*/

package main

import "fmt"

func main() {
	/* Here we utilise the Printf method that is part of the fmt package. The problem specifies that we use the %d format code, which is used to specify an integer, a base 10 int to be precise. */
	fmt.Printf("%d\n", 6)
}
