package main

import ( //importando pacotes
	"fmt"
	"time"
)

func f(from string) { //função f
	for i := 0; i < 3; i++ { //loop
		fmt.Println(from, ":", i) //imprimindo
	}
}

func main() { //função principal
	f("direct") //chamando a função f

	go f("goroutine") //chamando a função f em uma goroutine

	go func(msg string) { //chamando uma função anônima em uma goroutine
		fmt.Println(msg) //imprimindo
	}("going") //passando o parâmetro

	time.Sleep(time.Second) //pausando a execução por 1 segundo
	fmt.Println("done") //imprimindo

}