package main

import ( //Importando pacotes
	"fmt"
	"math/rand"
	"time"
)

func main() { //Função principal

	fmt.Print(rand.Intn(100), ",") //Gera um número aleatório entre 0 e 100
	fmt.Print(rand.Intn(100)) 	//Gera um número aleatório entre 0 e 100
	fmt.Println() //Quebra de linha

	fmt.Println(rand.Float64()) //Gera um número aleatório entre 0 e 1

	fmt.Print((rand.Float64()*5)+5, ",") //Gera um número aleatório entre 5 e 10
	fmt.Print((rand.Float64() * 5) + 5) //Gera um número aleatório entre 5 e 10
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano()) //Gera um número aleatório baseado no tempo
	r1 := rand.New(s1) //Gera um número aleatório baseado no tempo

	fmt.Print(r1.Intn(100), ",") //Gera um número aleatório entre 0 e 100
	fmt.Print(r1.Intn(100)) //Gera um número aleatório entre 0 e 100
	fmt.Println()

	s2 := rand.NewSource(42) //Gera um número aleatório baseado no tempo
	r2 := rand.New(s2) //Gera um número aleatório baseado no tempo
	fmt.Print(r2.Intn(100), ",") //Gera um número aleatório entre 0 e 100
	fmt.Print(r2.Intn(100)) //Gera um número aleatório entre 0 e 100
	fmt.Println() //Quebra de linha
	s3 := rand.NewSource(42) //Gera um número aleatório baseado no tempo
	r3 := rand.New(s3) //Gera um número aleatório baseado no tempo
	fmt.Print(r3.Intn(100), ",") //Gera um número aleatório entre 0 e 100
	fmt.Print(r3.Intn(100)) //Gera um número aleatório entre 0 e 100

}
