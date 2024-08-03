package main

import "fmt"

func main() {
	canal := make(chan string, 2) // tamanho do buffer do canal
	canal <- "Hello word"         // o canal só será fechado quando o canal atingir o tamanho do buffer
	canal <- "go lang "

	mensagem := <-canal
	fmt.Println(mensagem)
	mensagem2 := <-canal
	fmt.Println(mensagem2)
}
