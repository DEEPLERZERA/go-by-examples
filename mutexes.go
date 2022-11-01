package main

import ( //Importando pacotes
	"fmt"
	"sync"
)


type Container struct { //Criando uma struct
	mu     sync.Mutex //Mutex
	counters map[string]int //Mapa
}

  
func (c *Container) inc(name string) { //Função para incrementar
	c.mu.Lock() //Bloqueando o mutex
	defer c.mu.Unlock() //Desbloqueando o mutex
	c.counters[name]++ //Incrementando o valor
}
 
func main() { //Função principal
	c := Container { //Criando um objeto da struct

		counters: map[string]int{"a": 0, "b": 0}, //Inicializando o mapa

	}

	var wg sync.WaitGroup //Criando um WaitGroup

	doIncrement := func(name string, n int) { //Função para incrementar
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	
	wg.Add(3) //Adicionando 3 ao WaitGroup
	go doIncrement("a", 1000) //Chamando a função para incrementar
	go doIncrement("b", 1000)
	go doIncrement("c", 1000)
  
	wg.Wait() //Esperando o WaitGroup terminar
	fmt.Println(c.counters) //Imprimindo o mapa
}