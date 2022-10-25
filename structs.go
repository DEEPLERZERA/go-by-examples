package main

import "fmt"

type person struct { //Criando struct
	name string
	age int
}

func newPerson(name string) *person { //Função de criar nova pessoa
	p := person{name: name} //Criando variável
	p.age = 42
	return &p //Retornando ponteiro 
}

func main() { //Função principal
	fmt.Println(person{"Bob", 20}) //Imprimindo resultado

	fmt.Println(person{name: "Alice", age: 30}) 

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s //Criando ponteiro
	fmt.Println(sp.age) //Imprimindo resultado
 
	sp.age = 51 //Atribuindo valor 
	fmt.Println(sp.age) //Imprimindo resultado
}