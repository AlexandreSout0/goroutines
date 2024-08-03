package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonnacci(numero)
	}
}

/* Passo a posicao da sequencia de fibonnaci, e a func me devolve o valor dessa posicao */
func fibonnacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonnacci(posicao-2) + fibonnacci(posicao-1)
}
