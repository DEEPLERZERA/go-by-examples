package main

import "fmt"

func main() { // função principal

	messages := make(chan string) // cria um canal de string

	go func() { messages <- "ping" }() // envia "ping" para o canal

	msg  := <-messages // recebe do canal e atribui a msg
	fmt.Println(msg) // imprime msg
}