package main

import (  //Importando pacotes
	"fmt"
	"time"
)

func main() { //Função principal
	p := fmt.Println //Atribuindo a variável p a função Println do pacote fmt

	t := time.Now() //Pegando tempo 
	p(t.Format(time.RFC3339)) //Imprimindo o tempo no formato RFC3339 

	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")  //Atribuindo a variável t1 a função Parse do pacote time
	p(t1) //Imprimindo a variável t1

	p(t.Format("3:04PM")) //Imprimindo o tempo no formato 3:04PM
	p(t.Format("Mon Jan _2 15:04:05 2006")) //Imprimindo o tempo no formato Mon Jan _2 15:04:05 2006
	p(t.Format("2006-01-02T15:04:05.999999-07:00")) //Imprimindo o tempo no formato 2006-01-02T15:04:05.999999-07:00
	form := "3 04 PM" //Atribuindo a variável form ao formato 3 04 PM
	t2, e := time.Parse(form, "8 41 PM") //Convertendo o formato 3 04 PM para 8 41 PM
	p(t2) //Imprimindo a variável t2

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", //Imprimindo o tempo no formato 2006-01-02T15:04:05-00:00
		t.Year(), t.Month(), t.Day(), 
		t.Hour(), t.Minute(), t.Second()) 

	ansic := "Mon Jan _2 15:04:05 2006" //Atribuindo a variável ansic ao formato Mon Jan _2 15:04:05 2006
	_, e = time.Parse(ansic, "8:41PM") //Convertendo o formato 8:41PM para o formato Mon Jan _2 15:04:05 2006
	p(e) //Imprimindo a variável e
}