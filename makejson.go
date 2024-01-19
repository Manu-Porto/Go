package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*WWrite a program which prompts the user to first enter a name,
and then enter an address. Your program should create a map and
add the name and address to the map using the keys “name” and
“address”, respectively. Your program should use Marshal() to
create a JSON object from the map, and then your program should
print the JSON object.*/

func main(){
	
	var name, address string
	var mapa map[string]string

	reader := bufio.NewReader(os.Stdin) // é usado para criar um leitor de buffer para a entrada padrão

	fmt.Printf("Insert your name: ")
	name, _ = reader.ReadString('\n')  //é usado para ler a entrada até encontrar uma nova linha (Enter).
	fmt.Printf("Now, your address: ")
	address, _ = reader.ReadString('\n')  //é usado para ler a entrada até encontrar uma nova linha (Enter).  
	
	mapa = map [string]string {"name": strings.TrimSpace(name), "address": strings.TrimSpace(address)}
	barr, err := json.MarshalIndent(mapa, "", " ")

	if err != nil {
		fmt.Printf("Error number %s", err)
		os.Exit(1)
	}

	fmt.Println(string(barr))
}