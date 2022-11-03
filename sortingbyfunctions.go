package main


import ( // importando pacotes
	"fmt"
	"sort"
)

type byLength []string // criando um tipo de slice de strings

func (s byLength) Len() int { // implementando o método Len
	return len(s) // retornando o tamanho do slice
}

func (s byLength) Swap(i, j int) {  // implementando o método Swap
	s[i], s[j] = s[j], s[i] // trocando os valores de i e j
}

func (s byLength) Less(i, j int) bool { // implementando o método Less
	return len(s[i]) < len(s[j]) // retornando se o tamanho de i é menor que o tamanho de j
}

func main() { // função principal
	fruits := []string{"peach", "banana", "kiwi"} // criando um slice de strings
	sort.Sort(byLength(fruits)) // ordenando o slice de strings
	fmt.Println(fruits) // imprimindo o slice de strings
}

