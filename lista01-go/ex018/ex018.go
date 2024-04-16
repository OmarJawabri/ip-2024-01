package main

import "fmt"

func main() {
	var x, a1, r, n, r1, SomaNumeros int
	x = 2
	fmt.Scanln(&a1, &r, &n)
	SomaNumeros = a1
	r1 = r
	for n >= x {
		SomaNumeros = SomaNumeros + a1 + r
		r = r + r1
		x++
	}
	fmt.Println(SomaNumeros)
}
