package main

import (
	"fmt"
	"sort"
	"strconv"
)

func conjuntouniao(numeros1, numeros2 []int)[]int{
	uniao := []int{}
	for i := 0; i < len(numeros1); i++ {
		if i > 0 && numeros1[i] != 0 && numeros1[i] != numeros2[i]{
			uniao = append(uniao, numeros1[i])
		} 
		if i > 0 && numeros2[i] != 0 && numeros2[i] != numeros1[i]{
			uniao = append(uniao, numeros2[i])
		}
		if i > 0 && numeros2[i] != 0 && numeros1[i] != 0 && numeros1[i] == numeros2[i]{
			uniao = append(uniao, numeros1[i])
		}
		if i == 0 && numeros2[i] != numeros1[i]{
			uniao = append(uniao, numeros2[i])
			uniao = append(uniao, numeros1[i])
		}
		if i == 0 && numeros2[i] == numeros1[i]{
			uniao = append(uniao, numeros1[i])
		}
	}
	sort.Ints(uniao)
	return uniao
}

func conjuntointerseccao(numeros1, numeros2 []int)[]int{
	interseccao := []int{}
	for i := 0; i < len(numeros1); i++ {
		for k := 0; k < len(numeros2); k++ {
			if i > 0 && numeros1[i] != 0 && numeros1[i] == numeros2[k]{
				interseccao = append(interseccao, numeros1[i])
			}
			if i == 0 && numeros1[i] == numeros2[k]{
				interseccao = append(interseccao, numeros1[i])
			}
		}
	}
	return interseccao
}

func inteiroString(x []int)string{
	convertido := ""
	for i := 0; i < len(x); i++ {
		convertido += strconv.Itoa(x[i])
		if i != len(x)-1{
			convertido += ","
		}
	}
	return convertido
}

func main(){
	var t1, t2, y, x int
	fmt.Scanln(&t1)
	fmt.Scanln(&t2)
	for t1 < 1 || t1 > 100 || t2 < 1 || t2 > 100{
		fmt.Println("Valores excederam ou estão abaixo do limite, digite novamente:")
		fmt.Scanln(&t1)
		fmt.Scanln(&t2)
	}
	numeros1 := make([]int, 100)
	numeros2 := make([]int, 100)
	if t1 >= 1 && t1 <= 100 && t2 >= 1 && t2 <= 100{
		for i := 0; i < t1; i++ {
			fmt.Scanln(&numeros1[i])
			for k := 0; k < t1; k++ {
				if numeros1[i] == numeros1[k] && i != k{
					y++
				}
			}
			if y != 0{
				fmt.Println("Numero igual a outro, digite novamente: ")
				fmt.Scanln(&numeros1[i])
			}
			y = 0
		}
		for i := 0; i < t2; i++ {
			fmt.Scanln(&numeros2[i])
			for k := 0; k < t2; k++ {
				if numeros2[i] == numeros2[k] && i != k{
					x++
				}
			}
			if x != 0{
				fmt.Println("Numero igual a outro, digite novamente: ")
				fmt.Scanln(&numeros2[i])
			}
			x = 0
		}
		uniao := conjuntouniao(numeros1, numeros2)
		for i := 0; i < len(uniao); i++ {
			for k := 0; k < len(uniao); k++ {
				if uniao[i] == uniao[k] && i != k{
					uniao = append(uniao[:i], uniao[i+1:]...)
				}
			}
		}
		interseccao := conjuntointerseccao(numeros1, numeros2)
		convertido1 := inteiroString(uniao)
		convertido2 := inteiroString(interseccao)
		fmt.Printf("(%s)\n", convertido1)
		fmt.Printf("(%s)", convertido2)
	}
}