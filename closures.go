package main

import "fmt"

func intSeq() func() int { //Função que retorna uma função
	i := 0 //Variável local
	return func() int { //Função anônima
		i++ //Incrementando variável
		return i //Retornando variável
	}
}

func main() { //Função principal
	nextInt := intSeq() //Chamando função

	fmt.Println(nextInt()) //Imprimindo resultado
	fmt.Println(nextInt()) //Imprimindo resultado
	fmt.Println(nextInt()) //Imprimindo resultado

	newInts := intSeq() //Chamando função
	fmt.Println(newInts()) //Imprimindo resultado
}