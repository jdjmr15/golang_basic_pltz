package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			println(i)
		}
	}

	fmt.Printf("\n")

	var i int64 = 0
	strNum := ""
	for i < 10 {
		if i%2 == 1 {
			strNum = strconv.FormatInt(i, 32)
			fmt.Printf("%d:%T -> %T\n", i, i, strNum)
		}
		i++
	}

	fmt.Printf("\n")
	modulo := 5 % 2

	switch modulo {
	case 0:
		fmt.Println("Numero Par")
	default:
		fmt.Println("Numero Impar")
	}

}
