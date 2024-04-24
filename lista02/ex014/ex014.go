package main

import (
	"fmt"
)

func main(){
	var m, n, x int
	x = 1
	const numValor = 999
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	var linha[numValor]int
	for i := 2; i <= m; i++ {
		for k := 1; k <= n; k++ {
			if i > k {
				linha[x] = i
			}
			x++
		}
	}
	y :=1
	for i := 2; i <= m; i++ {
		for k := 1; k <= n; k++ {
			if i > k {
				if linha[y] == linha[y+1] {
					fmt.Printf("(%d, %d) - ", i, k)
				} else{
					fmt.Printf("(%d, %d)\n", i, k)
				}
			}
			y++
		}
	}
}
