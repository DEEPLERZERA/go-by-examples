package main

import "fmt"

func sum(nums ...int) { //Função com parâmetro variável
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums { //Iterando parâmetro
		total += num //Somando valores
	}
	fmt.Println(total) //Imprimindo resultado
}

func main() { //Função principal
	sum(1, 2) //Chamando função
	sum(1, 2, 3) //Chamando função

	
	nums := []int{1, 2, 3, 4} //Criando slice
	sum(nums...) //Chamando função
}