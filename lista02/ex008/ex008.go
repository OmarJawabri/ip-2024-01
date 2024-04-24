package main

import "fmt"

func main(){
	var n int
	k:= 1
	fmt.Scanln(&n)
	var time [99]int
	f := n - 2
	var x, y int = 1, 1
	for i := n; i > 1; i-- {
		x = x*i
	}
	for i := f; i > 1; i-- {
		y = y*i
	}
	l := x/2
	combinacao := l/y
	for i := 1; i <= n; i++ {
		time[i-1] = i
	}
	for i := 0; i <= int(combinacao); i++ {
		for n := 1; n <= int(combinacao); n++ {
			if time[i+n] != 0 {
				fmt.Printf("Final %d: Time%d X Time%d\n", k, time[i], time[i+n])
				k++
			}
		}
	}
	if n < 2 {
		fmt.Println("Campeonato inválido!")
	}
}