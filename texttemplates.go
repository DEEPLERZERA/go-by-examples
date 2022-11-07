package main

import ( //Importando pacotes
	"os"
	"text/template"
)

func main() { //Função principal
	t1 := template.New("t1") //Criando um novo template
	t1, err := t1.Parse("Value is {{.}}\n") //Parseando o template
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value is {{.}}\n")) //Parseando o template

	t1.Execute(os.Stdout, "Hello, World!") //Executando o template

	t1.Execute(os.Stdout, 123)

	t1.Execute(os.Stdout, []string {  
		"Go",
		"Rust",
		"Python",
		"JavaScript",
	})

	Create := func(name, t string) *template.Template { //Criando uma função
		return template.Must(template.New(name).Parse(t)) //Retornando o template
	}

	t2 := Create("t2", "Name: {{.Name}}\n") //Criando um novo template

	t2.Execute(os.Stdout, struct { //Executando o template
		Name string
	}{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string { //Executando o template
		"Name": "Mickey Mouse",
	})

	t3 := Create("t3", "{{if . -}}Yes{{else -}}No{{end -}}\n") //Criando um novo template

	t3.Execute(os.Stdout, "Not empty") //Executando o template
	t3.Execute(os.Stdout, "") //Executando o template

	t4 := Create("t4", "{{range . -}}{{.}} {{end -}}\n") //Criando um novo template

	t4.Execute(os.Stdout, []string { //Executando o template
		"Go",
		"Rust",
		"Python",
		"JavaScript",
	})

}