package main

import "fmt"

func main() {
	// Declaracion constante.
	const pi float64 = 3.14
	const pi2 = 3.1415

	// Variables
	base := 12
	var altura int = 14
	// Variable sin dato es igual 0
	var area int

	// Imprimir constante
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Variables
	fmt.Println(base)
	fmt.Println(altura)
	fmt.Println(area)

	// Area del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del cuadrado: ", areaCuadrado)

	// Paquete fmt
	fmt.Println("Paquete fmt")
	holaMensaje := "Hola"
	palabraMensaje := "Mundo"

	fmt.Println(holaMensaje, palabraMensaje)

	var nombre string = "Platzi"
	var cursos int16 = 500

	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprint
	mensaje := fmt.Sprintf("Sprintf: %s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Print(mensaje)

	// Tipo de datos
	fmt.Printf("Tipo de datos para la variable nombre: %T\n", nombre)
	fmt.Printf("Tipo de datos para la variable cursos: %T\n", cursos)
}
