package main

import (
	"fmt"
)

func main() {
	var nota1, nota2, nota3, MEDIA float64
	fmt.Scanln(&nota1, &nota2, &nota3)
	MEDIA = (nota1 + nota2 + nota3) / 3
	if MEDIA >= 6 {
		fmt.Printf("MÉDIA = %.2f\n", MEDIA )
		fmt.Println("APROVADO")
	} else if MEDIA < 6 {
		fmt.Printf("MÉDIA = %.2f\n", MEDIA)
		fmt.Println("REPROVADO")
	}
}
