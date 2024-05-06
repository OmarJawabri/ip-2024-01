package main

import (
	"fmt"
	"strconv"
)

func dividir(x int)[]int{
	quebrado := []int{}
	for x > 0{
		quebrado = append(quebrado, x%10)
		x = x / 10
	}
	for i, j := 0, len(quebrado)-1; i < j; i, j = i+1, j-1 {
		quebrado[i], quebrado[j] = quebrado[j], quebrado[i]
	}

	return quebrado
}

func apagar(x []int, tirar int)[]int{
	var retirar int = 10001
	menor := make([]int, tirar)
	for k := 0; k < tirar; k++ {
		menor[k] = 10001
		for i := 0; i < len(x); i++ {
			if menor[k] > x[i]{
				menor[k] = x[i]
			}
		}
		for i := 0; i < len(x); i++ {
			if menor[k] == x[i]{
				retirar = i
				if retirar != 10001{
					break
				}
			}
		}
		x = append(x[:retirar], x[retirar+1:]...)
	}
	return x
}

func inteiroString(x []int)string{
	convertido := ""
	for _, c:= range x{
		convertido += strconv.Itoa(c)
	}
	return convertido
}

func main(){
	var n, d, valor, tirar, y int
	n, d = 1, 1
	armazenamento := []string{}
	for n != 0 && d != 0{
		fmt.Scanln(&n, &d)
		tirar = n - d
		if d >= 1 && d <= 100000 && n >= 1 && n <= 100000{
			fmt.Scanln(&valor)
		}
		quebrar := dividir(valor)
		valorapagado := apagar(quebrar, tirar)
		convertido := inteiroString(valorapagado)
		armazenamento = append(armazenamento, convertido)
		y++
	}
	for i := 0; i < y-1; i++ {
		fmt.Println(armazenamento[i])
	}
}