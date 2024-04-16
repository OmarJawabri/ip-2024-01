package main

import (
	"fmt"
	"strconv"
)

func main(){
	var n1, n2, n3 string
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&n3)
	var z string = n1 + n2 + n3
	var a, b, c, i int
	a, _ = strconv.Atoi(n1)
	b, _ = strconv.Atoi(n2)
	c, _ = strconv.Atoi(n3)
	i, _ = strconv.Atoi(z)
	if a < 10 && b < 10 && c < 10{
		fmt.Print(z,", ",i*i)
	}else{
		fmt.Println("Digito Inválido!")
	}
}