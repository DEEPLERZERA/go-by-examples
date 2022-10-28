package main

import ( // pacotes necessários
	"fmt"
	"time"
)

func main() { //Função principal
	c1 := make(chan string, 1) // cria um canal de string
	go func() { // função anônima goroutine
		time.Sleep(time.Second * 2) // espera 2 segundos
		c1 <- "result 1" // envia "result 1" para o canal c1
	}()
 
	select { // select para cases
	case res := <-c1: // recebe do canal c1 e atribui a res
		fmt.Println(res)
	case <-time.After(time.Second * 1): // recebe do canal time.After e atribui a res
		fmt.Println("timeout 1")
}
	c2 := make(chan string, 1) // cria um canal de string
	go func() { // função anônima goroutine
		time.Sleep(time.Second * 2) // espera 2 segundos
		c2 <- "result 2"  // envia "result 2" para o canal c2
	}()

	select { // select para cases
	case res := <-c2: // recebe do canal c2 e atribui a res
		fmt.Println(res)  
	case <-time.After(time.Second * 3): // recebe do canal time.After e atribui a res
		fmt.Println("timeout 2") 
	}
}