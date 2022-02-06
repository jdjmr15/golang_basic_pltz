package main

import "fmt"

func normalFunction(mensaje string) {
	fmt.Println(mensaje)
}

func tripleArgumentos(a, b int, c string) {
	fmt.Println(a, b, c)
}

func retornarValor(z int) int {
	return z * 2

}

func retornarDosValores(z int) (a, b int) {
	return z, z * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgumentos(1, 2, "String")
	fmt.Println("Valor retornado: ", retornarValor(2))

	x, y := retornarDosValores(2)
	fmt.Printf("Valor 1: %d\nValor 2: %d\n", x, y)

	var w, _ int = retornarDosValores(2)
	fmt.Printf("Valor 1: %d\nValor 2: %d\n", w, 0)
}
