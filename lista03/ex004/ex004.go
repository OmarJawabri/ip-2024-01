package main

import "fmt"

func main(){
	var n int
	fmt.Scanln(&n)
	vetor := make([]int, n, n)
	if n < 5000 && n > 0{
		for i := 0; i < n; i++ {
			fmt.Scan(&vetor[i])
		}
		for _, c := range vetor{
			fmt.Printf("%d ", c)
		}
	}	
}