package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("Insert the text:")
	reader := bufio.NewReader(os.Stdin) // é usado para criar um leitor de buffer para a entrada padrão
	Text, _ := reader.ReadString('\n')  //é usado para ler a entrada até encontrar uma nova linha (Enter).

	te_xt := strings.ToLower(Text)
	text := strings.Replace(te_xt, " ", "", -1)

	if strings.HasPrefix(text, "i") && strings.Contains(text, "a") && strings.HasSuffix(strings.TrimSpace(text), "n") { //aqui foi difícil, o n não ia por nada
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not found!")
	}
}
