package main

import (
	"fmt"
	"math"
)

func main() {
	var Numero float64
	fmt.Scan(&Numero)
	if math.Mod(Numero, 3) == 0 && math.Mod(Numero, 5) == 0{
		fmt.Println("O NÚMERO É DIVISÍVEL")
	}else{
		fmt.Println("O NÚMERO NÃO É DIVISÍVEL")
	}
}
