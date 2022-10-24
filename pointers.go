package main

import "fmt"

func zeroval(ival int) { //Função com um parâmetro
	ival = 0 //Atribuindo valor
}

func zeroptr(iptr *int) { //Função com um ponteiro
	*iptr = 0 //Atribuindo valor
}

func main() { //Função principal
	i := 1 //Criando variável
	fmt.Println("initial:", i) //Imprimindo resultado

	zeroval(i) //Chamando função
	fmt.Println("zeroval:", i) //Imprimindo resultado

	zeroptr(&i) //Chamando função com variável com pointeiro
	fmt.Println("zeroptr:", i) //Imprimindo resultado

	fmt.Println("pointer:", &i) //Imprimindo resultado da variável com endereço de memória
}