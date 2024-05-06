package main

import "fmt"

func main(){
	const num = 9999
	var vetor[num]int
	var n int
	fmt.Scanln(&n)
	if n < 5000 && n > 0{
		for i := 0; i < n; i++ {
			fmt.Scan(&vetor[i])
		}
		for i := n-1; i >= 0; i-- {
			fmt.Print(vetor[i], " ")
		}
	}
}