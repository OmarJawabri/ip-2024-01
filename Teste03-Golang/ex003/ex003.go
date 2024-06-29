package main

import "fmt"

func bubbleSort(lista []int){
	n := len(lista)
	for i := 0; i < n-1; i++ {
		for j := n-1; j > i; j-- {
			if lista[j] < lista[j-1]{
				lista[j], lista[j-1] = lista[j-1], lista[j]
			}
		}
	}
}

func main(){
	array := []int{1, 45, 65, 32, 12, 24, 67, 1, 23}
	bubbleSort(array)
	fmt.Println(array)
}