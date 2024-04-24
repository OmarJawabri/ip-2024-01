package main

import (
	"fmt"
	"math"
)

func main(){
	const numVezes = 99999
	var n, e, rk float64
	var x int
	fmt.Scanln(&n)
	fmt.Scanln(&e)
	var r[numVezes]float64
	var erro[numVezes]float64
	rk = 1
	if n > 0{
		for k := 0; math.Abs(n - math.Pow(rk, 2)) > e ; k++ {
			rk = (rk + (n/rk))/2
			erro[k] = math.Abs(math.Pow(rk, 2) - n)
			r[k] = rk
			x++
		}
		for k := 0; x > k; k++ {
			fmt.Printf("r: %.9f, erro: %.9f\n", r[k], erro[k])
		}
	} else if n == 0{
		fmt.Printf("r: %.9f, erro: %.9f", n, e)
	} else if n < 0{
		fmt.Println("Número não possui raiz quadrada")
	}
}
