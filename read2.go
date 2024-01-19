package main 

/*Write a program which reads information from a file and represents it in a slice of structs. 
Assume that there is a text file which contains a series of names. Each line of the text file 
has a first name and a last name, in that order, separated by a single space on the line. 

Your program will define a name struct which has two fields, fname for the first name, and 
lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will 
successively read each line of the text file and create a struct which contains the first 
and last names found in the file. Each struct created will be added to a slice, and after all
lines have been read from the file, your program will have a slice containing one struct for 
each line in the file. After reading all lines from the file, your program should iterate 
through your slice of structs and print the first and last names found in each struct.

Submit your source code for the program, “read.go”.*/

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

var names []Name

func main () {

	// file name
	fmt.Printf("Insert the file name: ")
	var file string
	fmt.Scan(&file)

	// openning file
	f, err := os.Open(file)
	if (err != nil){
		fmt.Printf("Erro number %s", err)
		os.Exit(1)
	}

	// saving
	scanner := bufio.NewScanner(f) //idk why using the "New"

	// reading each line
	for scanner.Scan() {
		line:= scanner.Text ()

			// spliting the name and last name
			Split := strings.Fields(line)

			// join again
			names = append (names, Name{fname : Split[0], lname : Split [1]})
	}
		
	// checking erros
	if err := scanner.Err(); err!= nil {
		fmt.Printf("Erro number %s", err)
		os.Exit(1)
	}

	// ok? if so, we will print
	for _, n := range names {
		fmt.Printf("Name: %s | Last: %s \n", n.fname, n.lname)
	}
	
}

