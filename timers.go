package main 


import ( //Importando pacotes
	"fmt"
	"time"
)

func main() { //Função principal
	
	timer1 :=  time.NewTimer(2 * time.Second) //Criando um timer com 2 segundos de duração

	<- timer1.C //Aguardando o timer acabar
	fmt.Println("Timer 1 expired") //Imprimindo mensagem
 
	timer2 := time.NewTimer(time.Second) //Criando um timer com 1 segundo de duração
	go func() { //Criando uma goroutine
		<- timer2.C //Aguardando o timer acabar
		fmt.Println("Timer 2 expired") //Imprimindo mensagem
	}() 
	stop2 := timer2.Stop() //Parando o timer
	if stop2 { //Verificando se o timer foi parado
		fmt.Println("Timer 2 stopped") //Imprimindo mensagem
	}

	time.Sleep(2 * time.Second) //Aguardando 2 segundos
}