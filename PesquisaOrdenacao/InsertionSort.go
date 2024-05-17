package main

import "fmt"

func insertionSort(lista []int) {
	for j := 1; j < len(lista); j++ {
		chave := lista[j]
		i := j - 1
		for i >= 0 && lista[i] > chave {
			lista[i+1] = lista[i]
			i--
		}
		lista[i+1] = chave
	}
}

/*	EM ALGORITMO
	InsertionSort(lista)
	Seja n o tamanho de lista
	Para j = 1 ate n
		chave <- lista[j]
		i <- j-1
		Enquanto i >= 0 e lista[i] > chave
			lista[i+1] <- lista[i]
			i--
		lista[i+1] <- chave
*/

func main() {
	lista1 := []int{10, 23, 43, 32, 21, 2}
	insertionSort(lista1)
	fmt.Println(lista1)
}