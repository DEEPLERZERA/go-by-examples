package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {  // Criando a função MapKeys
	keys := make([]K, 0, len(m)) // Criando um slice de K com o tamanho e capacidade do map m
	for k := range m {  // Iterando sobre o map m
		keys = append(keys, k) // Adicionando a chave k ao slice keys
	}
	return keys // Retornando o slice de K
}


type List[T any] struct { // Criando um tipo List[T] que é um slice de T
	head, tail *element[T] // Criando dois ponteiros para o tipo element[T]
} 


type element[T any] struct { // Criando um tipo element[T] que é um elemento de uma lista
	next *element[T]
	val T
}


func (lst *List[T]) Push(v T) { // Criando um método Push para o tipo List[T]
	if lst.tail == nil { // Se o ponteiro tail for nil
		lst.head = &element[T]{val: v} // Atribuindo um ponteiro para um elemento com o valor v ao ponteiro head
		lst.tail = lst.head
	} else { // Se o ponteiro tail não for nil
		lst.tail.next = &element[T]{val: v} // Atribuindo um ponteiro para um elemento com o valor v ao ponteiro next do elemento tail
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T { // Criando um método GetAll para o tipo List[T]
	var elems []T // Criando um slice de T
	for e := lst.head; e != nil; e = e.next { // Iterando sobre a lista
		elems = append(elems, e.val)
	}
	return elems
}

func main() { // Criando uma instância da struct List[T]
	var m = map[int]string{1: "2", 2: "4", 4: "8"} // Criando um map de int para string

	fmt.Println("keys:", MapKeys(m)) // Imprimindo as chaves do map m

	_ = MapKeys[int, string](m) // Chamando a função MapKeys com os tipos int e string

	lst := List[int]{} // Criando uma instância da struct List[int]
	lst.Push(10) // Chamando o método Push da struct List[int]
	lst.Push(20)
	lst.Push(30)
	fmt.Println("list:", lst.GetAll()) // Imprimindo o slice de int da struct List[int]

}