/*
Problem #9: Using a For Loop
Author: Asam Shan
Date Created: 21/05/2019
Comments:
*/

/*
***************** Using a For Loop *****************
Using your editor create a file called testdata9 that contains 5 positive integer values. Write a  program that
reads and prints all five values. Each number should be read using the fscanf function embedded inside a for-loop.
You should not use an array. You should a for-loop.
*/

package main

import (
	"log"
	"os" // used for opening a file
)

func main() {

	// fnuction to open an file using the os package
	file, err := os.Open("testdata9") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	// Todo add rest
}
