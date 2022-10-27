package main

import "fmt"

func main() { // função principal

	messages := make(chan string, 2) // cria um canal de string com buffer de 2

	messages <- "buffered" // envia "buffered" para o canal
	messages <- "channel" // envia "channel" para o canal

	fmt.Println(<-messages) // recebe do canal e imprime
	fmt.Println(<-messages)

}