package main

import "fmt"

func main() {
	m := make(map[string]int)   //Criando mapa

	m["k1"] = 7 			   //Adicionando valores
	m["k2"] = 13

	fmt.Println("map:", m) 	//Imprimindo mapa
 
	v1 := m["k1"] 			//Acessando valores
	fmt.Println("v1: ", v1) //Imprimindo valor

	fmt.Println("len:", len(m)) //Imprimindo tamanho do mapa

	delete(m, "k2") //Deletando valor
	fmt.Println("map:", m) //Imprimindo mapa

	_, prs := m["k2"] //Verificando se valor existe
	fmt.Println("prs:", prs) //Imprimindo resultado

	n := map[string]int{"foo": 1, "bar": 2} //Criando mapa com valores
	fmt.Println("map:", n) //Imprimindo mapa
}

