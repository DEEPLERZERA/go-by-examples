package main

import (  //Importando pacotes
    "bytes"
    "fmt"
    "regexp"
)

func main() { //Função principal
 
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach") //Verifica se a string contém o padrão
    fmt.Println(match) //Imprime o resultado

    r, _ := regexp.Compile("p([a-z]+)ch") //Compila o padrão

    fmt.Println(r.MatchString("peach")) //Verifica se a string contém o padrão

    fmt.Println(r.FindString("peach punch")) //Encontra a primeira ocorrência do padrão

    fmt.Println("idx:", r.FindStringIndex("peach punch")) //Encontra a primeira ocorrência do padrão e retorna o índice

    fmt.Println(r.FindStringSubmatch("peach punch")) //Encontra a primeira ocorrência do padrão e retorna o índice e o padrão

    fmt.Println(r.FindStringSubmatchIndex("peach punch")) //Encontra a primeira ocorrência do padrão e retorna o índice e o padrão

    fmt.Println(r.FindAllString("peach punch pinch", -1)) //Encontra todas as ocorrências do padrão

    fmt.Println("all:", r.FindAllStringSubmatchIndex(  //Encontra todas as ocorrências do padrão e retorna o índice e o padrão
        "peach punch pinch", -1))

    fmt.Println(r.FindAllString("peach punch pinch", 2)) //Encontra as duas primeiras ocorrências do padrão

    fmt.Println(r.Match([]byte("peach"))) //Verifica se a string contém o padrão

    r = regexp.MustCompile("p([a-z]+)ch") //Compila o padrão
    fmt.Println("regexp:", r)

    fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) //Substitui o padrão por outra string

    in := []byte("a peach") //Substitui o padrão por outra string
    out := r.ReplaceAllFunc(in, bytes.ToUpper) //Substitui o padrão por outra string e transformar em letra maiuscula
    fmt.Println(string(out)) //Imprime na tela
}