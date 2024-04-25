package main

import (
	"fmt"
	"strconv"
)

func stringInversa(s string)string{
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = 1+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


func main(){
	var T int
	numero1 := make([]int, 999)
	numero2 := make([]int, 999)
	fmt.Scanln(&T)
	for i := 0; i < T; i++ {
		fmt.Scanln(&numero1[i], &numero2[i])
	}
	for i := 0; i < T; i++ {
		x1 := strconv.Itoa(numero1[i])
		y1 := strconv.Itoa(numero2[i])
		inversox := stringInversa(x1)
		inversoy := stringInversa(y1)
		inverso1, _ := strconv.Atoi(inversox)
		inverso2, _ := strconv.Atoi(inversoy)
		if inverso1 > inverso2{
			fmt.Println(inverso1)
		} else{
			fmt.Println(inverso2)
		}
	}
}