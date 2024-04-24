package main

import (
	"fmt"
	"math"
)

func main(){
	var y, x, codigoVenda, codigoLucro, menor10, maior10, maior20, maisVenda  int
	var maiorLucro, compraT, vendaT, lucroT float64
	x = 1
	const numVezes = 999
	var codigo[numVezes]int
	var compra[numVezes]float64
	var venda[numVezes]float64
	var numVenda[numVezes]int
	var lucro[numVezes]float64
	for i := 0; x != 0; i++ {
		fmt.Scanln(&codigo[i], &compra[i], &venda[i], &numVenda[i])
		x = codigo[i]
		y++
	}
	for i := 0; i < y; i++ {
		lucro[i] = (venda[i] - compra[i])/compra[i]
		if  lucro[i] < 0.1{
			menor10 = menor10 + 1
		} else if lucro[i] >= 0.1 && lucro[i] <= 0.2{
			maior10 = maior10 + 1
		} else if lucro[i] > 0.2{
			maior20 = maior20 + 1
		}
		if lucro[i] > maiorLucro{
			maiorLucro = lucro[i]
			codigoLucro = codigo[i]
		}
		if numVenda[i] > maisVenda{
			maisVenda = numVenda[i]
			codigoVenda = codigo[i]
		}
		compraT = compraT + compra[i]*float64(numVenda[i])
		vendaT = vendaT + venda[i]*float64(numVenda[i])
		lucroT = ((vendaT - compraT)/compraT)*100
		if math.IsNaN(lucroT) {
			lucroT = 0
		}
	}
		fmt.Println("Quantidade de mercadorias que geraram lucro menor que 10%: ", menor10)
		fmt.Println("Quantidade de mercadorias que geraram lucro maior ou igual a 10% e menor ou igual a 20%: ", maior10)
		fmt.Println("Quantidade de mercadorias que geraram lucro maior do que 20%: ", maior20)
		fmt.Printf("Código da mercadoria que gerou maior lucro: %d\n", codigoLucro)
		fmt.Printf("Código da mercadoria mais vendida: %d\n",codigoVenda)
		fmt.Printf("Valor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f", compraT, vendaT, lucroT)
		fmt.Println("%")
}