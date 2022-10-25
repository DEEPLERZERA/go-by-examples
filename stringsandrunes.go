package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "Hello, 世界" //A string literal, just like a rune literal, is a constant.


	fmt.Println(len(s)) // 13

	for i := 0; i < len(s); i++ { // 1 byte at a time
		fmt.Printf("%x ", s[i]) 
	}

	fmt.Println() 

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 9


	for idx, runeValue := range s { // 4 bytes at a time
		fmt.Printf("%#U starts at %d\n", runeValue, idx) 
	}

	fmt.Println("\nUsing DecodeRuneInString") 
	for i, w := 0, 0; i < len(s); i += w { // 4 bytes at a time
		runeValue, width := utf8.DecodeRuneInString(s[i:]) // DecodeRuneInString returns the rune and its width in bytes
		fmt.Printf("%#U starts at %d\n", runeValue, i) 
		w = width

		examineRune(runeValue) //Examine the rune
	}


}

func examineRune(r rune) { //Function to examine the rune

	if r == 't'{ //If the rune is equal to 't'
		fmt.Println("Found a tee") //Print the message
	} else if r == '世' { //If the rune is equal to '世'
		fmt.Println("Found so sua") //Print the message

	}
	
}
