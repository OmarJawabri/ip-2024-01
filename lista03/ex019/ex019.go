package main

import (
	"fmt"
	"sort"
)

func inteirosUnicos(x []int)[]int{
	unico := 1
	for i := 1; i < len(x); i++ {
		if x[i-1] != x[i]{
			x[unico] = x[i]
			unico++
		}
	}
	return x[:unico]
}

func main(){
	var N int
	fmt.Scanln(&N)
	inteiros := make([]int, N)
	if N >= 1 && N <= 1000{
		for i := 0; i < N; i++ {
			fmt.Scanln(&inteiros[i])
		}
		sort.Ints(inteiros)
		intunico := inteirosUnicos(inteiros)
		for i := 0; i < len(intunico); i++ {
			fmt.Println(intunico[i])
		}
	}
}