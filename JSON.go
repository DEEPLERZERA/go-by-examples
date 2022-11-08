package main

import ( //Importando pacotes
	"encoding/json"
	"fmt"
	"os"
)


type response1 struct { //Criando uma struct
	Page   int
	Fruits []string
}

type response2 struct { //Criando uma struct
	Page   int      `json:"page"` //Atribuindo um nome para o campo json
	Fruits []string `json:"fruits"`
}

func main() { //Função principal

	bolB, _ := json.Marshal(true) //Converte um valor para json
	fmt.Println(string(bolB))   //Imprime na tela

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"} //Criando um slice
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{ //Criando um ponteiro
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`) //Criando um slice

	var dat map[string]interface{} //Criando um map

	if err := json.Unmarshal(byt, &dat); err != nil { //Deserializa o json
		panic(err) //Imprime na tela
	}

	fmt.Println(dat)

	num := dat["num"].(float64) //Converte o valor para float64
	fmt.Println(num)

	strs := dat["strs"].([]interface{}) //Converte o valor para interface
	str1 := strs[0].(string) //Converte o valor para string
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}` //Criando uma string json
	res := response2{} //Criando uma struct
	json.Unmarshal([]byte(str), &res) //Deserializa o json
	fmt.Println(res) //Imprime na tela
	fmt.Println(res.Fruits[0]) //Imprime na tela com o indice 0	

	enc := json.NewEncoder(os.Stdout) //Cria um novo encoder
	d := map[string]int{"apple": 5, "lettuce": 7} //Cria um map
	enc.Encode(d) //Codifica o json
}