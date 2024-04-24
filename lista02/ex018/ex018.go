package main

import "fmt"

func main(){
	var n1, n2, n3, maiorV, MMC int
	MMC = 1
	fmt.Scanln(&n1, &n2, &n3)
	if n1 >= n2 && n1 >= n3 {
		maiorV = n1
	} else if n2 >= n1 && n2 >= n3{
		maiorV = n2
	} else if n3 >= n1 && n3 >= n2{
		maiorV = n3
	}
	if n1 > 0 && n2 > 0 && n3 > 0{
		for i := 2; i <= maiorV; i++ {
			for k := 0; k < maiorV; k++ {
				if n1 % i == 0 || n2 % i == 0 || n3 % i == 0 {
					fmt.Printf("%d %d %d :%d\n", n1, n2, n3, i)
					MMC = MMC*i
				}
				if n1 % i == 0{
					n1 = n1/i
				}
				if n2 % i == 0{
					n2 = n2/i
				}
				if n3 % i == 0{
					n3 = n3/i
				}
				if n1 % i != 0 && n2 % i != 0 && n3 % i != 0 {
					k++
				}
			}
		}
	}
	if n1 > 0 && n2 > 0 && n3 > 0{
		fmt.Printf("MMC: %d", MMC)
	}
}

