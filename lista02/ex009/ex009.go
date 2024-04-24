package main

import (
	"fmt"
)

func main(){
	var n int
	var naoPrimo string
	fmt.Scanln(&n)
	for i := 2; i < n; i++ {
		if n % i == 0 {
			naoPrimo = "NÃO PRIMO"
		} 
	}
	if n <= 1 {
		fmt.Println("NÃO PRIMO")
	} else if naoPrimo == "NÃO PRIMO" {
		fmt.Println(naoPrimo)
	} else{
		fmt.Println("PRIMO")
	}
}