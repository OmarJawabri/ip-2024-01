package main

import (
	"fmt"
)

func main(){
	var T int
	fmt.Scanln(&T)
	numero1 := make([]int, 999)
	numero2 := make([]int, 999)
	inverso := make([]int, 999)
	inverso2 := make([]int, 999)
	for i := 0; i < T; i++ {
		fmt.Scanln(&numero1[i], &numero2[i])
	}
	for i := 0; i < T; i++ {
		for numero1[i] > 0{
			x := numero1[i] % 10
			inverso[i] = inverso[i]*10 + x
			numero1[i] = numero1[i]/10
		}
		numero1[i] = inverso[i]
	}
	for i := 0; i < T; i++ {
		for numero2[i] > 0 {
			x := numero2[i] % 10
			inverso2[i] = inverso2[i]*10 + x
			numero2[i] = numero2[i]/10
		}
		numero2[i] = inverso2[i]
	}
	for i := 0; i < T; i++ {
		if numero1[i] > numero2[i] {
			fmt.Println(numero1[i])
		} else{
			fmt.Println(numero2[i])
		}
	}
}