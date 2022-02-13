package main

import "fmt"

func main() {
	diccionario := make(map[string]int)
	diccionario["Jose"] = 14
	diccionario["Pepito"] = 20
	
	for k,v := range diccionario {
		fmt.Println(k,v)
	}	

	// Valor existe
	value, ok := diccionario["Jose"]
	fmt.Println(value,ok)
	value, ok = diccionario["Jesus"]
	fmt.Println(value,ok)
}
