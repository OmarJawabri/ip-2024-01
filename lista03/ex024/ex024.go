package main

import (
	"fmt"
	"strconv"
	"os"
)

func preencherVCount(vcont, vetor []int, N int)[]int{
	for i := 0; i < N; i++{
		for k := 0; k < len(vcont); k++ {
			if vetor[i] == k{
				vcont[k]++
			}
		}
	}
	for i := 1; i < len(vcont); i++ {
		vcont[i] += vcont[i-1]
	}
	return vcont
}

func preencherVOrd(vcont, vetor, vord []int)[]int{
	for i := len(vetor)-1; i >= 0; i-- {
		for k := 0; k < len(vcont); k++ {
			if vetor[i] == k{
				vord[vcont[k]] = vetor[i]
				vcont[k] = vcont[k] -1
			}
		}
	}
	return vord
}

func inteiroString(x []int, N int)string{
	armazenar := ""
	for i := 0; i < len(x); i++ {
		if i != 0 && i <= N{
		if i != 1 && i != N+1{
			armazenar += " "
		}
		armazenar += strconv.Itoa(x[i])
		}
	}
	return armazenar
}

func main(){
	var N, maior int
	N = 1
	armazenamento := []string{}
	for N != 0{
		maior = 0
		fmt.Scanln(&N)
		if N > 1 && N <= 10000{
			vetor := make([]int, N)
			for i := 0; i < N; i++ {
				fmt.Scanln(&vetor[i])
				if vetor[i] < 0 || vetor[i] > 10000{
					os.Exit(0)
				}
			}
			for i := 0; i < N; i++ {
				if maior < vetor[i]{
					maior = vetor[i]
				}
			}
			vOrd := make([]int, maior+2)
			vCount := make([]int, maior+1)
			vCount = preencherVCount(vCount, vetor, N)
			vOrd = preencherVOrd(vCount, vetor, vOrd)
			stringss := inteiroString(vOrd, N)
			armazenamento = append(armazenamento, stringss)
			vetor = nil
			vCount = nil
			vOrd = nil
		}
	}
	for i := 0; i < len(armazenamento); i++ {
		fmt.Println(armazenamento[i])
	}
}