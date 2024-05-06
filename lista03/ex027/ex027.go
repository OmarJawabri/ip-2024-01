package main

import (
	"fmt"
	"os"
)

func intercalados(v1, v2, vr []int)[]int{
	for i := 0; i < len(vr); i++ {
		if v1[i] < v2[i]{
			vr[i] = v1[i]
			if v2[len(vr)-1] == 0{
				v2 = append([]int{0}, v2...)
			}
		} else{
			vr[i] = v2[i]
			if v1[len(vr)-1] == 0{
				v1 = append([]int{0}, v1...)
			}
		}
		if i > 2 && v1[i] == 0{
			vr[i] = v2[i]
		} else if i > 2 && v2[i] == 0{
			vr[i] = v1[i]
		}
	}
	return vr
}

func main(){
	var q1, q2 int
	fmt.Scanln(&q1)
	fmt.Scanln(&q2)
	v1 := make([]int, 10000)
	v2 := make([]int, 10000)
	vr := make([]int, q1+q2)
	if q1 > 0 && q1 <= 500000 && q2 > 0 && q2 <= 500000{
		for i := 0; i < q1; i++ {
			fmt.Scanln(&v1[i])
			if v1[i] < 0 || v1[i] > 999999{
				os.Exit(0)
			}
		}
		for k := 0; k < q2; k++ {
			fmt.Scanln(&v2[k])
			if v2[k] < 0 || v2[k] > 999999{
				os.Exit(0)
			}
		}
		for i := 1; i < q1; i++ {
			if v1[i] < v1[i-1]{
				os.Exit(0)
			}
		}
		for i := 1; i < q2; i++ {
			if v2[i] < v2[i-1]{
				os.Exit(0)
			}
		}
		vr = intercalados(v1, v2, vr)
		for i := 0; i < len(vr); i++ {
			fmt.Println(vr[i])
		}
	}
}