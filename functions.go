package main

import "fmt"

func plus (a int, b int) int { //Função com dois parâmetros e um retorno
	return a + b
}

func plusPlus (a, b, c int) int { //Função com três parâmetros e um retorno
	return a + b + c
}

func main() { //Função principal
	res := plus(1, 2) //Chamando função
	fmt.Println("1+2 =", res) //Imprimindo resultado

	res = plusPlus(1, 2, 3) //Chamando função
	fmt.Println("1+2+3 =", res) //Imprimindo resultado
}