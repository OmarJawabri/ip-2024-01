package main

import (
	"fmt"
	"math"
)

func main() {
	var n, y int32
	var x float64
	n = 0
	fmt.Scanln(&x, &y)
	if math.Mod(x, 2) == 0 {
		for y > n{
		fmt.Printf("%v ", x)
			x = x + 2
			n++
		}
	} else{
		fmt.Println("O Primeiro Número não é par!")
	}
}
