package main

import (
	"fmt"
	"math"
)

func main() {
	var salariominimo, kwgastos, custoporkw, custodoconsumo, custocomdesconto float64
	fmt.Scan(&salariominimo)
	fmt.Scan(&kwgastos)
	custoporkw = salariominimo * 0.007
	fmt.Println("Custo por kW: R$", math.Round(custoporkw*100)/100)
	custodoconsumo = kwgastos*custoporkw
	fmt.Println("Custo do consumo: R$", math.Round(custodoconsumo*100)/100)
	custocomdesconto = custodoconsumo*0.9
	fmt.Println("Custo com desconto: R$", math.Round(custocomdesconto*100)/100)
}