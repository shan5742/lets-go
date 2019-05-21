/*
Problem #3: The Character P
Author: Asam Shan
Date Created: 21/05/2019
*/

/*
***************** The Character P *****************
Write a program that prints out the single character P. I want you to use the %c format code.
*/

package main

import "fmt"

func main() {
	/* Very similar to both problem 1 and 2, this time we use Printf combined with the %c format specifier, which as per the docs is "the character represented by the corresponding Unicode code point". I did have some issues here, simply writing "%c, P" returns undefined as the %c prints out the character that is associated with the given int, in the case P = 80 */
	fmt.Printf("%c\n", 80)
}
