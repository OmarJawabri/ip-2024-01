package main

import (
	"fmt"
	"math"
)

func main() {
	var AlturaP, ArestaH, Volume, Ab float64
	fmt.Scan(&AlturaP)
	fmt.Scan(&ArestaH)
	Ab = (3 * math.Pow(ArestaH, 2) * math.Sqrt(3))/2
	Volume = Ab*AlturaP/3
	fmt.Printf("O VOLUME DA PIRÂMIDE É = %.2f METROS CÚBICOS", Volume)
}
