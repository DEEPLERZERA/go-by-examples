package main

import "fmt" //Importando pacotes

func mayPanic() { //Função que pode gerar um panic
	panic("a problem") //Gerando um panic
}

func main() { //Função principal

	defer func() { //Função anônima
		if r := recover(); r != nil { //Verificando se o panic foi gerado
			fmt.Println("Recovered. Error:\n", r) //Imprime a mensagem
		}
	}()

	mayPanic() //Chamando a função que pode gerar um panic
	fmt.Println("After panic") //Imprime a mensagem
}