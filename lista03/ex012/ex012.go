package main

import (
	"fmt"
	"sort"
)

func main(){
	var N int
	fmt.Scanln(&N)
		if N >= 1 && N <= 1000{
		inteiros := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scanln(&inteiros[i])
		}
		sort.Ints(inteiros)
		for i := 0; i < N; i++ {
			fmt.Println(inteiros[i])
		}
	}
}