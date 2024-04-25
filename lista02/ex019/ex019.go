package main

import "fmt"

func main(){
	const numValores = 999
	var n, i, x, l, y int
	y = 2
	x = 2
	i = 1
	var multiplicacao[numValores]int
	multiplicacao[1] = 1
	fmt.Scanln(&n)
	if n > 0 {
		fmt.Printf("%d*%d*%d = %d\n", i, i, i, multiplicacao[1])
		for k := 1; k < n; k++ {
			i++
			fmt.Printf("%d*%d*%d = ", i, i, i)
			for h := 1; h <= i; h++{
				multiplicacao[x] = multiplicacao[x-1] + 2
				l = multiplicacao[x]
				x++
			}
			for h := 1; h <= i; h++ {
				multiplicacao[y] = multiplicacao[y-1] + 2
				fmt.Printf("%d", multiplicacao[y])
				if multiplicacao[y] != l {
					fmt.Printf("+")
				}
				y++
			}
			fmt.Printf("\n")
			
		}
	}
}