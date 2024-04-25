package main

import "fmt"

func main(){
	const numValor = 999
	var sim, k, x int
	x = 1
	var ordem[numValor]string
	var numero[numValor]float64
	var n[numValor]int
	for k = 0; x != 0; k++ {
		sim = 0
		fmt.Scan(&n[k])
		x = n[k]
		for i := 1; i <= x; i++ {
			fmt.Scan(&numero[i])
		}
		for i := 1; i <= x; i++ {
			if numero[i-1] < numero[i]{ //VER SE <= CONTA TAMBÉM
				sim++
			}
		}
		if x != 0{
			if sim == x {
				ordem[k] = "ORDENADA"
			} else{
				ordem[k] = "DESORDENADA"
			}
		}
	}
	for i := 0; i <= k; i++ {
		fmt.Println(ordem[i])
	}
}