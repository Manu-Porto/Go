package main

import (
	"fmt"
)

/*Write a program which prompts the user to enter integers and stores the integers
in a sorted slice. The program should be written as a loop. Before entering the loop,
the program should create an empty integer slice of size (length) 3. During each pass
through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the
slice in sorted order. The slice must grow in size to accommodate any number of integers
which the user decides to enter. The program should only quit (exiting the loop) when the
user enters the character ‘X’ instead of an integer.*/

func slice() {

	sli := make([]string, 0, 3)
	var leitura string
	fmt.Println("Hi, friend!")

	for {
		fmt.Printf("\nInsert a number: ")
		fmt.Scan(&leitura)

		/* aqui é para ler se é X*/
		if (leitura == "X") || (leitura == "x") {
			fmt.Printf("No integrer")
			break
		}

		/*x começa com 0 e cresce a medida que informações são colocadas*/
		sli = append(sli, leitura)
		fmt.Println(sli)

	}
}
