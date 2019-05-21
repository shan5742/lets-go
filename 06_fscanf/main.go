/*
Problem #6: The fscanf function
Author: Asam Shan
Date Created: 21/05/2019
Comments: Source for the solution https://golang.org/pkg/os/
*/

/*
***************** The fcsanf function *****************
Using your favourite editor, create file called testdata6 that has a single integer value stored into it. Write a program
that opens a file named testdata6, reads the single value stored in it using the fscanf function, and then prints the
value back onto the screen. You should call the function fclose before you return to the Unix operating system. Read
about the function fclose for a better understanding of it.
*/

package main

import (
	"fmt"
	"log"
	"os" // used for opening a file
)

func main() {

	// fnuction to open an file using the os package
	file, err := os.Open("testdata6") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	//we can now read the data, store the number of bytes as well as the count, which is the number within the testdata file

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes and the number in testdata6 is: %q\n", count, data[:count])
}
