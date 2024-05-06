package main

import (
	"fmt"
	"os"
)

func checagem(x []int, n1, n2, n3, n4, n5, n6 int)string{
	var y int
	var resultado string
	for i := 0; i < len(x); i++ {
		if x[i] == n1 || x[i] == n2 || x[i] == n3 || x[i] == n4 || x[i] == n5 || x[i] == n6{
			y++
		}
	}
	if y == 4{
		resultado = "QUADRA"
	} else if y == 5{
		resultado = "QUINA"
	} else if y == 6{
		resultado = "SENA"
	} else{
		resultado = ""
	}
	return resultado
}

func main(){
	var n1, n2, n3, n4, n5, n6, N, q, qu, s int
	armazenamento := []string{}
	fmt.Scanln(&n1, &n2, &n3, &n4, &n5, &n6)
	fmt.Scanln(&N)
	if N >= 1 && N <= 50000{
		for i := 0; i < N; i++ {
			numeros := make([]int, 6)
			for k := 0; k < 6; k++ {
				fmt.Scan(&numeros[k])
				if numeros[k] < 1 || numeros[k] > 60{
					os.Exit(0)
				}
			}
			for k := 0; k < 6; k++ {
				for i := 0; i < 6; i++ {
					if numeros[k] == numeros[i] && k != i{
						os.Exit(0)
					}
				}
			}
			resultado := checagem(numeros, n1, n2, n3, n4, n5, n6)
			armazenamento = append(armazenamento, resultado)
			numeros = nil
		}
		for i := 0; i < len(armazenamento); i++ {
			if armazenamento[i] == "QUADRA"{
				q++
			} else if armazenamento[i] == "QUINA"{
				qu++
			} else if armazenamento[i] == "SENA"{
				s++
			}
		}
		if s != 0{
			fmt.Printf("Houve %d acertador(es) da sena\n", s)
		} else{
			fmt.Printf("Não houve acertador para sena\n")
		}
		if qu != 0{
			fmt.Printf("Houve %d acertador(es) da quina\n", qu)
		} else{
			fmt.Printf("Não houve acertador para quina\n")
		}
		if q != 0{
			fmt.Printf("Houve %d acertador(es) da quadra", q)
		} else{
			fmt.Printf("Não houve acertador para quadra")
		}
		
	}
}