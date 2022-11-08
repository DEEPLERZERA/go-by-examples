package main

import ( //Importando pacotes
    "encoding/xml"
    "fmt"
)

type Plant struct { //Criando uma struct
    XMLName xml.Name `xml:"plant"` //Atribuindo um nome para o campo xml
    Id      int      `xml:"id,attr"`
    Name    string   `xml:"name"`
    Origin  []string `xml:"origin"`
}

func (p Plant) String() string {  //Criando um método
    return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", //Imprime na tela
        p.Id, p.Name, p.Origin)
}

func main() {  //Função principal
    coffee := &Plant{Id: 27, Name: "Coffee"} //Criando um ponteiro
    coffee.Origin = []string{"Ethiopia", "Brazil"} //Atribuindo valores ao slice

    out, _ := xml.MarshalIndent(coffee, " ", "  ")  //Converte um valor para xml
    fmt.Println(string(out)) //Imprime na tela

    fmt.Println(xml.Header + string(out)) //Imprime na tela

    var p Plant  //Criando uma struct
    if err := xml.Unmarshal(out, &p); err != nil {  //Deserializa o xml
        panic(err) //Imprime na tela
    }
    fmt.Println(p)

    tomato := &Plant{Id: 81, Name: "Tomato"} //Criando um ponteiro
    tomato.Origin = []string{"Mexico", "California"} //Atribuindo valores ao slice

    type Nesting struct {  //Criando uma struct
        XMLName xml.Name `xml:"nesting"` //Atribuindo um nome para o campo xml
        Plants  []*Plant `xml:"parent>child>plant"`
    }

    nesting := &Nesting{} //Criando um ponteiro
    nesting.Plants = []*Plant{coffee, tomato} //Atribuindo valores ao slice

    out, _ = xml.MarshalIndent(nesting, " ", "  ") //Converte um valor para xml
    fmt.Println(string(out)) 	//Imprime na tela
}