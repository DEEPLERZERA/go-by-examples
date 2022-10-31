package main

import ( // importando pacotes
	"fmt"
	"time"
)

func main() { // função principal

	ticker := time.NewTicker(500 * time.Millisecond) // criando um ticker
	done := make(chan bool) 						// criando um canal

	go func() { // criando uma goroutine
		for { // loop infinito
			select {  // selecionando
			case <-done: // caso o canal done seja recebido
				return // retorna
			case t := <-ticker.C: // caso o ticker seja recebido
				fmt.Println("Tick at", t) // imprime
			}
		}
	}()
		time.Sleep(1600 * time.Millisecond) // pausa a execução por 1600 milisegundos
		ticker.Stop() // para o ticker
		done <- true // envia um valor para o canal done
		fmt.Println("Ticker stopped") // imprime
}
