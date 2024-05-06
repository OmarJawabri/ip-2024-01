package main

import (
	"fmt"
)

func conversorBinario(x int){
	var codigo[]int
	var y int
	for x != 0{
		z := x%2
		codigo = append(codigo, z)
		x = x/2
		y++
	}
	if len(codigo) == 0{
		fmt.Println("0")
	} else{
		for i := y-1; i >= 0; i-- {
			fmt.Printf("%d",codigo[i])
		}
		fmt.Println()
	}
	
}

func main(){
	var x, y, h int
	var numero[999]int
	for i := 0; x != -1; i++ {
		fmt.Scanln(&numero[i])
		x = numero[i]
		y++
	}
	for i := 0; i < y-1; i++ {
		h = numero[i]
		if h < 4294967295{
			conversorBinario(h)
		} else{
			fmt.Println("Número excede 32 bits")
		}
	}
}