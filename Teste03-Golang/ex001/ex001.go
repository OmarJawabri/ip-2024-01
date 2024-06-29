package main

import (
	"strings"
	"fmt"
	"os"	
	"bufio"
)

func main(){
	var resp string
	for resp != "#"{
		leitor := bufio.NewReader(os.Stdin)
		resp, _ := leitor.ReadString('\n')
		resp = strings.TrimSpace(resp)
		if resp == "#"{
			break
		}
		quebrar := strings.Split(resp, "")
		for i := len(quebrar)-1; i > 0; i-- {
			fmt.Print(quebrar[i])
		}
	}
}