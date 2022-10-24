package main

import "fmt"

func fact(n int) int { //Função recursiva
	if n == 0 { //Caso base
		return 1
	}
	return n * fact(n-1) //Chamada recursiva
}

func main() { //Função principal
	fmt.Println(fact(7)) //Imprimindo resultado
	
	var fib func(n int) int 

	fib = func(n int) int { //Função recursiva
		if n < 2 { //Caso base
			return n //Retornando valor
		}

		return fib(n-1) + fib(n-2) //Chamada recursiva
}

	fmt.Println(fib(7)) //Imprimindo resultado
}
