package main

import "fmt"

func main(){
	const numVezes = 99999
	var n, x, l int
	fmt.Scanln(&n)
	var div[numVezes]int
	var amigo[numVezes]int
	var amigo2[numVezes]int
	if n < 9 {
		for i := 1; i < 20000; i++ {
			for h := 1; h < i; h++ {
				if  i % h == 0 && i != h{
					div[i] = div[i] + h
				}
			}
		}
		for i := 1; i < 20000; i++ {
			for h := 1; h < i; h++ {
				if div[h] == i && div[i] == h{
					x++
				}
				if div[h] == i && div[i] == h{
					amigo[x] = h
					amigo2[x] = i
				}
			}
		}
		for i := 1; i < 20000; i++ {
			if amigo[i] != 0 && amigo2[i] != 0 {
				if l < n{
				fmt.Printf("(%d, %d)\n", amigo[i], amigo2[i])
				l++
				}
			}
		}	
	}
}