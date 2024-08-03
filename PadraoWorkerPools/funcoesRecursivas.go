package main

import "fmt"

/* Passo a posicao da sequencia de fibonnaci, e a func me devolve o valor dessa posicao */
func fibonnaci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonnaci(posicao-2) + fibonnaci(posicao-1)
}

/* Numeros da Sequencia: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765 */

func main2() {
	fmt.Println("Finonnaci Recursivamente")
	posicao := uint(45) // quantidade de numeros

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonnaci(i))
	}

}
