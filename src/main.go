package main

import "fmt"

func main() {
	//defer (Cierra conexcion y archivo)
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	i := 0
	for i < 10 {

		if i == 2 {
			i++
			continue
		}

		if i == 8 {
			break
		}
		fmt.Println(i)
		i++
	}
}
