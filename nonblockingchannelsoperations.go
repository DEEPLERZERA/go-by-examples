package main

import "fmt" //Importando fmt

func main() { //Função principal 
	messages := make(chan string) //Criando um canal de string
	signals := make(chan bool) //Criando um canal de boolean

	select { //Select
	case msg := <-messages: //Recebendo mensagem do canal messages
		fmt.Println("received message", msg) //Imprimindo mensagem recebida
	default: //Caso não receba mensagem
		fmt.Println("no message received") //Imprimindo mensagem
	}

	msg := "hi" //Mensagem
	select { //Select
	case messages <- msg: //Enviando mensagem para o canal messages
		fmt.Println("sent message", msg) //Imprimindo mensagem enviada
	default: //Caso não envie mensagem
		fmt.Println("no message sent") //Imprimindo mensagem
	}

	select { //Select 
	case msg := <-messages: //Recebendo mensagem do canal messages
		fmt.Println("received message", msg) //Imprimindo mensagem recebida
	case sig := <-signals: //Recebendo sinal do canal signals
		fmt.Println("received signal", sig) //Imprimindo sinal recebido
	default: //Caso não receba mensagem ou sinal
		fmt.Println("no activity") //Imprimindo mensagem
	}
}
