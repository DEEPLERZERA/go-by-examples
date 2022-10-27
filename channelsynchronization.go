package main

import ( // importação de pacotes
	"fmt"
	"time"
)

func worker(done chan bool) { // função worker
	fmt.Print("working...")  // imprime working...
	time.Sleep(time.Second) // dorme por 1 segundo 
	fmt.Println("done") // imprime done

	done <- true // envia true para o canal done
}

func main() { // função principal
	done := make(chan bool, 1) // cria um canal de bool com buffer de 1
	go worker(done) // inicia a função worker em uma goroutine

	<-done // recebe do canal done e ignora o valor
}
