package main

import ( //Importando pacotes
	"fmt"
	"time"
)

func main() { //Função principal
	p := fmt.Println //Criando um ponteiro

	now := time.Now() //Criando uma variável
	p(now) //Imprime na tela

	then := time.Date( //Criando uma variável
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC) //Atribuindo valores
	p(then)

	p(then.Year()) //Imprime na ano
	p(then.Month()) //Imprime na mês
	p(then.Day())  //Imprime na dia
	p(then.Hour()) //Imprime na hora
	p(then.Minute()) //Imprime na minuto
	p(then.Second()) //Imprime na segundo
	p(then.Nanosecond()) //Imprime na nanosegundo
	p(then.Location()) //Imprime na localização

	p(then.Weekday()) //Imprime na dia da semana

	p(then.Before(now))  //Imprime na tela
	p(then.After(now)) //Imprime na tela
	p(then.Equal(now)) //Imprime na tela

	diff := now.Sub(then) //Criando uma variável de diferença
	p(diff) //Imprime na tela
  
	p(diff.Hours()) //Imprime diferença de horas
	p(diff.Minutes()) //Imprime diferença de minutos
	p(diff.Seconds()) //Imprime diferença de segundos
	p(diff.Nanoseconds()) //Imprime diferença de nanosegundos
 
	p(then.Add(diff)) //Imprime na tela
	p(then.Add(-diff)) //Imprime na tela

}