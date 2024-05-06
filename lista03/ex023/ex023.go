package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func contarVogais1(s1 string)[]int{
	vogais := make([]int, 5)
	for _, c := range s1{
		switch c{
		case 'a', 'A':
			vogais[0]++
		case 'e', 'E':
			vogais[1]++
		case 'i', 'I':
			vogais[2]++
		case 'o', 'O':
			vogais[3]++
		case 'u', 'U':
			vogais[4]++
		}
	}	
	return vogais	
}

func contarVogais2(s2 string)[]int{
	vogais2 := make([]int, 5)
	for _, c2 := range s2{
		switch c2{
		case 'a', 'A':
			vogais2[0]++
		case 'e', 'E':
			vogais2[1]++
		case 'i', 'I':
			vogais2[2]++
		case 'o', 'O':
			vogais2[3]++
		case 'u', 'U':
			vogais2[4]++
		}
	}
	return vogais2
}

func main(){
	leitor := bufio.NewReader(os.Stdin)
	s1, _ := leitor.ReadString('\n')
	s1 = strings.TrimSpace(s1)
	if strings.Contains(s1, ";") == false{
		fmt.Println("FORMATO INVÁLIDO!")
	} else{
		separadas := strings.Split(s1, ";")
		s1 = separadas[0]
		s2 := separadas[1]
		if len(separadas[0]) > 1000 || len(separadas[1]) > 1000{
			fmt.Println("STRING EXCEDEU 1000 CARACTERES")
		}
		if len(separadas[0]) > 1000 || len(separadas[1]) > 1000{
			os.Exit(0)
		}	
		vogais := contarVogais1(s1)
		vogais2 := contarVogais2(s2)
		d := math.Pow(float64(vogais[0] - vogais2[0]), 2) + math.Pow(float64(vogais[1] - vogais2[1]), 2) + math.Pow(float64(vogais[2] - vogais2[2]), 2) + math.Pow(float64(vogais[3] - vogais2[3]), 2) + math.Pow(float64(vogais[4] - vogais2[4]), 2)
		distancia := math.Sqrt(d)
		fmt.Printf("(%d, %d, %d, %d, %d)\n", vogais[0], vogais[1], vogais[2], vogais[3], vogais[4])
		fmt.Printf("(%d, %d, %d, %d, %d)\n", vogais2[0], vogais2[1], vogais2[2], vogais2[3], vogais2[4])
		fmt.Printf("%.2f", distancia)
	}
}