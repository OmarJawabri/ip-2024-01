package main

import "fmt"

func validar1(x []int)string{
	var soma, digito1, y int
	var validar1 string
	y = 1
	for i := 0; i < 9; i++ {
		soma = soma + x[i]*y
		y++
	}
	if soma % 11 != 10{
		digito1 = soma % 11
	} else{
		digito1 = 0
	}
	if digito1 == x[9] {
		validar1 = "VALIDO"
	} else{
		validar1 = "INVALIDO"
	}
	return validar1
}

func validar2(x []int)string{
	var h, soma, digito2 int
	var validar2 string
	h = 9
	for i := 0; i < 9; i++ {
		soma = soma + x[i]*h
		h--
	}
	if soma % 11 != 10{
		digito2 = soma % 11
	} else{
		digito2 = 0
	}
	if digito2 == x[10]{
		validar2 = "VALIDO"
	} else{
		validar2 = "INVALIDO"
	}
	return validar2
}

func main(){
	var T int
	var valido1, valido2 string
	fmt.Scanln(&T)
	valido := make([]string, T)
	for i := 0; i < T; i++ {
		cpf := make([]int, 11, 11)
		for k := 0; k < 11; k++ {
			fmt.Scan(&cpf[k])
		}
		valido1	= validar1(cpf)
		valido2 = validar2(cpf)
		if valido1 == "VALIDO" && valido2 == "VALIDO"{
			valido[i] = "CPF válido"
		} else{
			valido[i] = "CPF inválido"
		}
		cpf = nil
	}
	for i := 0; i < T; i++ {
		fmt.Println(valido[i])
	}
}