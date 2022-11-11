package main

import ( //Importando pacotes
	"fmt"
	"strconv"
)

func main() { //Função principal
  
	p, _ := strconv.ParseBool("true") //Converte string para boolean
	fmt.Println("ParseBool:", p) //Imprime o resultado

	f, _ := strconv.ParseFloat("1.234", 64) //Converte string para float
	fmt.Println("ParseFloat:", f) //Imprime o resultado

	i, _ := strconv.ParseInt("-123", 0, 64) //Converte string para int
	fmt.Println("ParseInt:", i) //Imprime o resultado

	d, _ := strconv.ParseUint("789", 0, 64) //Converte string para uint
	fmt.Println("ParseUint:", d) //Imprime o resultado

	k, _ := strconv.Atoi("135") //Converte string para int
	fmt.Println("Atoi:", k)  //Imprime o resultado

	_, e := strconv.Atoi("wat") //Converte string para int
	fmt.Println("Atoi:", e) //Imprime o resultado
}