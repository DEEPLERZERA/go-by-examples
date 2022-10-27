package main

import "fmt"

func ping(pings chan<- string, msg string) { // função ping
	pings <- msg // envia msg para o canal pings
}

func pong(pings <-chan string, pongs chan<- string) { // função pong
	msg := <-pings // recebe do canal pings e atribui a msg
	pongs <- msg // envia msg para o canal pongs
}

func main() { // função principal
	pings := make(chan string, 1) // cria um canal de string com buffer de 1
	pongs := make(chan string, 1) // cria um canal de string com buffer de 1
	ping(pings, "passed message") // chama a função ping
	pong(pings, pongs) // chama a função pong
	fmt.Println(<-pongs) // recebe do canal pongs e imprime
}