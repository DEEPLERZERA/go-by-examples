package main

import ( // pacotes necessários
	"fmt"
	"time"
)

func main() { //Função principal

	c1 := make(chan string) // cria um canal de string
	c2 := make(chan string)

	go func() { // função anônima goroutine
		time.Sleep(time.Second * 1) // espera 1 segundo
		c1 <- "one" // envia "one" para o canal c1
	}()

	go func() { // função anônima goroutine
		time.Sleep(time.Second * 2)
		c2 <- "two" // envia "two" para o canal c2
	}()

	for i := 0; i < 2; i++ { // laço de repetição
		select { // select para cases
		case msg1 := <-c1: // recebe do canal c1 e atribui a msg1
			fmt.Println("received", msg1)  // imprime msg1
		case msg2 := <-c2: // recebe do canal c2 e atribui a msg2
			fmt.Println("received", msg2) // imprime msg2
		}
	}


}