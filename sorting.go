package main

import ( // importando pacotes
	"fmt"
	"sort"
)

func main() { // função principal
	strs := []string{"c", "a", "b"} // criando um slice de strings
	sort.Strings(strs) 			// ordenando o slice de strings
	fmt.Println("Strings:", strs) 	// imprimindo o slice de strings

	ints := []int{7, 2, 4} // criando um slice de inteiros
	sort.Ints(ints)  // ordenando o slice de inteiros
	fmt.Println("Ints:   ", ints) // imprimindo o slice de inteiros

	s := sort.IntsAreSorted(ints) // verificando se o slice de inteiros está ordenado
	fmt.Println("Sorted: ", s) // imprimindo o resultado da verificação
}