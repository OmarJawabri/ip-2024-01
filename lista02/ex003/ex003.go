package main

import "fmt"

func main(){
	var n, fat float64
	fmt.Scanln(&n)
	fat = 1 
	i:= n
	for i > 1{
		fat = fat*i
		i = i - 1
	}
	fmt.Printf("%.0f! = %.0f", n, fat)
}