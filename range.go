package main

import "fmt"

func main() {
	nums := []int{2, 3, 4} //Criando slice
	sum := 0 //Criando variável
	for _, num := range nums { //Iterando slice
		sum += num //Somando valores
	}

	fmt.Println("sum:", sum) //Imprimindo resultado

	for i, num := range nums { //Iterando slice
		if num == 3 {
			fmt.Println("index:", i) //Imprimindo index
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"} //Criando mapa
	for k, v := range kvs { //Iterando mapa
		fmt.Printf("%s -> %s", k, v) //Imprimindo valores
	} 

	for k := range kvs { //Iterando mapa
		fmt.Println("key:", k)
	}

	for i, c := range "go" { //Iterando string
		fmt.Println(i, c)
	}

}