package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) { // criando uma função f1 que retorna um inteiro e um erro
	if arg == 42 { // se o argumento for igual a 42
		return -1, errors.New("can't work with 42") // retorna -1 e um erro
	}
	return arg + 3, nil // retorna o argumento + 3 e nil
}

type argError struct { // criando uma struct argError
	arg  int
	prob string
}

func (e *argError) Error() string { // criando um método Error para a struct argError
	return fmt.Sprintf("%d - %s", e.arg, e.prob) // retorna uma string formatada
}

func f2(arg int) (int, error) { // criando uma função f2 que retorna um inteiro e um erro
	if arg == 42 { // se o argumento for igual a 42
		return -1, &argError{arg, "can't work with it"} // retorna -1 e um erro
	}
	return arg + 3, nil 		// retorna o argumento + 3 e nil
}
 
func main() { // criando a função main

	for _, i := range []int{2, 42} { // criando um loop que percorre um slice de inteiros
		if r, e := f1(i); e != nil { // se o retorno da função f1 for diferente de nil
			fmt.Println("f1 failed:", e)  // imprime a mensagem f1 failed: e o erro
		} else { // se não
			fmt.Println("f1 worked:", r) // imprime a mensagem f1 worked: e o retorno da função f1
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42) // atribuindo o retorno da função f2 para a variável e
	if ae, ok := e.(*argError); ok { // se o tipo da variável e for igual a argError
		fmt.Println(ae.arg) // imprime o valor do campo arg da struct argError
		fmt.Println(ae.prob) // imprime o valor do campo prob da struct argError
	}
}