package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go myFuncWrite("Hello Word", canal)
	fmt.Println("Depois da func myFuncWrite() ter sido executada ")

	mensagem := <-canal // Esperando o canal receber uma variavel (ou valor)
	fmt.Println(mensagem)
	fmt.Println("depois do canl receber a variavel enviada!")

	// Recebe mensagens e se o canal foi fechado ele sai do laço
	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	// Mesma coisa que o for de cima , mas usando range

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("fim de tudo!")

}

func myFuncWrite(text string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- text //enviando variavel (ou valor) para dentro do canal
		time.Sleep(time.Second)
	}
	close(canal) //fecha o canal
}
