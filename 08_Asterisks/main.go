/*
Problem #8: One Horizontal Line of Asterisks
Author: Asam Shan
Date Created: 21/05/2019
Comments:
*/

/*
***************** One Horizontal Line of Asterisks *****************

Using your preferred editor, create a file called testdata8 that has a single integer value stored into it. The integer will be a
number between 1 and 30 inclusive. Create a program that uses the fscanf function to read in this integer value and
then prints a horizontal line containing that many asterisksÕ (*). Use a for-loop in your program. For example,

If the file contained the integer 8 your program would print out:        ********
If the file contained the integer 15 your program would print out:      ***************


*/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//Similar to problem 6 in terms of using the os package to read a file and grab the data.
	// first set some variables to store data

	var number, i int

	// fnuction to open an file using the os package
	file, err := os.Open("testdata8") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	//we can now read the data and store the number from file in our variable we created earlier

	_, scanErr := fmt.Fscanf(file, "%d", &number)
	if scanErr != nil {
		fmt.Println(scanErr)
		return
	}

	// for loop to count
	for i = 0; i < number; i++ {
		fmt.Printf("*")
	}
	fmt.Printf("\n")
	// closing file
	file.Close()
}
