package main

import ( //Importando pacotes
	"fmt"
	"time" 
)


func main() { //Função principal
	
	requests := make(chan int, 5) //Criando um canal de inteiros com buffer de 5 posições
	for i := 1; i <= 5; i++ { //Loop para enviar 5 valores para o canal
		requests <- i //Enviando valores para o canal
	}
	close(requests) //Fechando o canal

	limiter := time.Tick(200 * time.Millisecond) //Criando um ticker que irá enviar um valor a cada 200 milisegundos


	for req := range requests { //Loop para receber os valores do canal
		<-limiter //Recebendo os valores do ticker
		fmt.Println("request", req, time.Now()) //Imprimindo o valor recebido do canal e a hora atual
	}

	burstyLimiter := make(chan time.Time, 3) //Criando um canal de tempo com buffer de 3 posições

	for i := 0; i < 3; i++ { //Loop para enviar 3 valores para o canal
		burstyLimiter <- time.Now() //Enviando valores para o canal
	}

	go func() { //Criando uma goroutine
		for t := range time.Tick(200 * time.Millisecond) { //Loop para enviar valores para o canal a cada 200 milisegundos
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5) //Criando um canal de inteiros com buffer de 5 posições
	for i := 1; i <= 5; i++ { //Loop para enviar 5 valores para o canal
		burstyRequests <- i //Enviando valores para o canal
	}
	close(burstyRequests) //Fechando o canal
	for req := range burstyRequests { //Loop para receber os valores do canal
		<-burstyLimiter //Recebendo os valores do canal
		fmt.Println("request", req, time.Now()) //Imprimindo o valor recebido do canal e a hora atual
	}

}