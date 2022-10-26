package main

import "fmt"

type base struct { // criando uma struct base
	num int
}

func (b base) describe() string { // criando um método para a struct base
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct { // criando uma struct container
	base
	str string
}

func main() { // criando uma instância da struct container
	co := container{base: base{num: 1,}, str: "hello",} // criando uma instância da struct container
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str) // acessando os campos da struct container

	fmt.Println("also num:", co.base.num) // acessando os campos da struct base

	fmt.Println("describe:", co.describe())

	type describer interface { // criando uma interface describer
		describe() string
	}

	var d describer = co // criando uma variável d do tipo describer e atribuindo a ela a instância da struct container
	fmt.Println("describer:", d.describe()) // acessando o método describe da interface describer
}
