package main

import (
	"fmt"
	"math"
)

func main(){
	var n int
	var hipotenusa, cateto, cateto2, h, l, x float64
	fmt.Scanln(&n)
	l, h, x = 1, 1, 1
	for i := 0; i < n*n*n; i++ {
		cateto = l
		cateto2 = h
		hipotenusa = math.Pow(x, 2)
		catetos := math.Pow(cateto, 2) + math.Pow(cateto2, 2)
		if catetos == hipotenusa && cateto <= cateto2 && cateto != 0{
			fmt.Printf("hipotenusa = %.0f, catetos %.0f e %.0f\n", x, cateto, cateto2)
		}
		if cateto2 == float64(n) {
			h = 0
			l++
		}
		if cateto == float64(n) && cateto2 == float64(n){
			x++
			l = 1
		}
		h++
	}
}