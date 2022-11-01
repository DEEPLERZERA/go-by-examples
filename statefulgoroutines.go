package main

import ( //Importando pacotes
	"fmt"
	"sync/atomic"
	"time"
	"math/rand"
)

type readOp struct { //Criando uma struct
	key  int
	resp chan int
}

type writeOp struct { //Criando uma struct
	key  int
	val  int	
	resp chan bool
}

func main() { //Função principal

	var readOps uint64  //Criando uma variável
	var writeOps uint64

	reads := make(chan readOp)  //Criando um canal
	writes := make(chan writeOp)

	go func() { //Criando uma função anônima
		var state = make(map[int]int) //Criando um map
		for { //Criando um loop infinito
			select { //Criando um select
			case read := <-reads: //Criando um case
				read.resp <- state[read.key] //Lendo o valor do map
			case write := <-writes:  //Criando um case
				state[write.key] = write.val //Escrevendo no map
				write.resp <- true //Escrevendo no canal
			}
		}
	}()

	for r := 0; r < 100; r++ { //Criando um laço de repetição
		go func() { //Criando uma função anônima
			for {
				read := readOp{ //Criando uma struct
					key:  rand.Intn(5), //Gerando um número aleatório
					resp: make(chan int)} //Criando um canal
				reads <- read //Escrevendo no canal
				<-read.resp //Lendo o canal
				atomic.AddUint64(&readOps, 1) //Adicionando 1 na variável
				time.Sleep(time.Millisecond) //Pausando a execução
			}
		}()
	}

	for w := 0; w < 10; w++ { 		//Criando um laço de repetição
		go func() { //Criando uma função anônima
			for {
				write := writeOp{ //Criando uma struct
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write //Escrevendo no canal
				<-write.resp
				atomic.AddUint64(&writeOps, 1) //Adicionando 1 na variável
				time.Sleep(time.Millisecond) //Pausando a execução
			}
		}()
	}

	time.Sleep(time.Second) //Pausando a execução

	readOpsFinal := atomic.LoadUint64(&readOps) //Carregando o valor da variável
	fmt.Println("readOps:", readOpsFinal) //Imprimindo o valor da variável
	writeOpsFinal := atomic.LoadUint64(&writeOps) //Carregando o valor da variável
	fmt.Println("writeOps:", writeOpsFinal) //Imprimindo o valor da variável
}