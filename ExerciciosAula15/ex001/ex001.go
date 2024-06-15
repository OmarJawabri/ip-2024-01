package main

import "fmt"

func elevado(x, n int) int{
	if n <= 0{
		return 1
	}
	return x * elevado(x, n-1)
}

func main(){
	var numero, elevar int
	fmt.Println("Digite o número e o valor que deseja elevar ele")
	fmt.Scanln(&numero)
	fmt.Scanln(&elevar)
	fmt.Println(elevado(numero, elevar))
}