package main

import "fmt"

func vals() (int, int) { //Função com dois retornos
	return 3, 7
}


func main() { //Função principal
	a, b := vals() //Chamando função
	fmt.Println(a) //Imprimindo resultado
	fmt.Println(b) //Imprimindo resultado

	_, c := vals() //Chamando função
	fmt.Println(c) //Imprimindo resultado
}