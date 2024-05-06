package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func separar7(x []int)[]int{
	separar := []int{}
	var somarest, somatot int
	for i := 0; i < len(x); i++ {
		for y := 0; y < len(x); y++ {
			somarest = x[i] + x[y]
			for k := 0; k < len(x); k++ {
				somatot = somatot + x[k]
			}
			if somatot - somarest == 100{
				for len(separar) < 7{
					for k := 0; k < len(x); k++ {
						if x[k] != x[i] && x[k] != x[y]{
							separar = append(separar, x[k])
						}
					}
				}
			}
			somatot = 0
			somarest = 0
		}
	}
	sort.Ints(separar)
	return separar
}

func inteiroString(x []int)string{
	convertido := ""
	for i := 0; i < len(x); i++ {
		convertido += strconv.Itoa(x[i])
		convertido += ","
	}
	return convertido
}

func main(){
	var T int
	fmt.Scanln(&T)
	var convertido string
	for i := 0; i < T; i++ {
		anoes := make([]int, 9)
		for k := 0; k < 9; k++ {
			fmt.Scanln(&anoes[k])
			if anoes[k] < 1 || anoes[k] > 99{
				os.Exit(0)
			}
		}
		for k := 0; k < 9; k++ {
			for i := 0; i < 9; i++ {
				if anoes[k] == anoes[i] && k != i{
					os.Exit(0)
				}
			}
		}
		separados := separar7(anoes)
		convertido += inteiroString(separados)
		anoes = nil
	}
	armazenar := strings.Split(convertido, ",")
	for i := 0; i < len(armazenar)-1; i++ {
		fmt.Println(armazenar[i])
	}
}