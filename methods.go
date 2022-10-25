package main

import "fmt"

type rect struct { //Criando struct
	 width, height int //Retornando valor
}

func (r *rect) area() int { //Função de calcular area com ponteiro
	return r.width * r.height
}

func (r rect) perim() int { //Função de calcular perimetro 
	return 2*r.width + 2*r.height //Retornando valor
}

func main() { //Função principal
	r := rect{width: 10, height: 5} //Criando variável

	fmt.Println("area: ", r.area()) //Imprimindo resultado
	fmt.Println("perim:", r.perim())
 
	rp := &r //Criando ponteiro
	fmt.Println("area: ", rp.area()) //Imprimindo resultado
	fmt.Println("perim:", rp.perim())
}