package main

import ( //Importando pacotes
	"fmt"
	"time"
)

func main() { //Função principal

	now := time.Now()  //Atribuindo a variável now a função Now() do pacote time
	fmt.Println(now)  //Imprimindo a variável now

	fmt.Println(now.Unix()) //Imprimindo a função Unix() da variável now
	fmt.Println(now.UnixMilli()) //Imprimindo a função UnixMilli() da variável now
	fmt.Println(now.UnixNano()) //Imprimindo a função UnixNano() da variável now

	fmt.Println(time.Unix(now.Unix(), 0) ) //Imprimindo a função Unix() da variável now
	fmt.Println(time.Unix(0, now.UnixNano())) //Imprimindo a função UnixNano() da variável now




}