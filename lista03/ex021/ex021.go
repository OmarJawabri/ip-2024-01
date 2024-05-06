package main

import (
	"fmt"
	"sort"
)

func pares(x []int){
	vpares := []int{}
	for i := 0; i < len(x); i++ {
		if x[i] % 2 == 0{
			vpares = append(vpares, x[i])
		}
	}
	sort.Ints(vpares)
	for i := 0; i < len(vpares); i++ {
		fmt.Println(vpares[i])
	}
}

func impares(x []int){
	vimpares := []int{}
	for i := 0; i < len(x); i++ {
		if x[i] % 2 != 0{
			vimpares = append(vimpares, x[i])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(vimpares)))
	for i := 0; i < len(vimpares); i++ {
		fmt.Println(vimpares[i])
	}
}

func main(){
	var N, y int
	fmt.Scanln(&N)
	inteiros := make([]int, N)
	if N > 1 && N < 100000{
		for i := 0; i < N; i++ {
			fmt.Scanln(&inteiros[i])
			if inteiros[i] < 0{
				y++
			}
		}
		if y == 0{
			pares(inteiros)
			impares(inteiros)
		}
	}
}