package main

import (
	"fmt"
)

func main() {
	var SalarioFuncionario, Reajuste float64
	fmt.Scan(&SalarioFuncionario)
	if SalarioFuncionario < 300 {
		Reajuste = SalarioFuncionario * 1.5
	} else {
		Reajuste = SalarioFuncionario * 1.3
	}
	fmt.Printf("SALÁRIO COM REAJUSTE = %.2f", Reajuste)
}
