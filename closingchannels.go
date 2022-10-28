package main

import "fmt" //Importa o pacote fmt

func main() { // Inicio da função main
	jobs := make(chan int, 5) // Cria um canal com buffer de 5 posições
	done := make(chan bool)

	go func() { // Inicio da função anônima
		for { // Inicio do loop infinito
			j, more := <-jobs // Recebe o valor do canal jobs
			if more { // Se o valor for verdadeiro
				fmt.Println("received job", j) // Imprime a mensagem
			} else { // Se o valor for falso
				fmt.Println("received all jobs") // Imprime a mensagem
				done <- true // Envia o valor true para o canal done
				return // Encerra a função
			}
		}
	}()

	for j := 1; j <= 3; j++ { // Inicio do loop
		jobs <- j // Envia o valor j para o canal jobs
		fmt.Println("sent job", j) // Imprime a mensagem
	}
	close(jobs) // Fecha o canal jobs
	fmt.Println("sent all jobs") // Imprime a mensagem

	<-done // Recebe o valor do canal done
}