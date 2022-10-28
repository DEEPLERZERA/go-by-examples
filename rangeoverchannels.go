package main

import "fmt" //Importa o pacote fmt
 
func main() { // Inicio da função main

	queue := make(chan string, 2) // Cria um canal com buffer de 2 posições
	queue <- "one"  // Envia o valor "one" para o canal queue
	queue <- "two" // Envia o valor "two" para o canal queue
	close(queue) // Fecha o canal queue

	for elem := range queue { // Inicio do loop percorrendo canal
		fmt.Println(elem)  // Imprime o valor do canal
	}
}