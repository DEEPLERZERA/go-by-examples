package main

import ( //Importando pacotes
	"fmt"	
	"sync"
	"sync/atomic"
)

func main() { //Função principal
	var ops uint64 //Variável do tipo uint64

	var wg sync.WaitGroup //Variável do tipo WaitGroup

	for i := 0; i < 50; i++ { //Loop de 50 vezes
		wg.Add(1) //Adicionando 1 ao WaitGroup

		go func() { //Função anônima
			for c := 0; c < 1000; c++ { //Loop de 1000 vezes
				
				
				atomic.AddUint64(&ops, 1) //Adicionando 1 ao valor da variável ops
			}
			wg.Done() //Removendo 1 do WaitGroup
		}()
	}

	wg.Wait() //Aguardando o WaitGroup zerar
	fmt.Println("ops:", ops) //Imprimindo o valor da variável ops
}