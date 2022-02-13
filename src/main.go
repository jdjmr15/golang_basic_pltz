package main

// Se deactivo <go env -w GO111MODULE=off>
import (
	"fmt"
	auto "mipaquete"
)

/*
Se move a la carpeta mipaquete
type auto struct {
	marca  string
	anio   int
	color  string
	modelo string
}

func main() {

	fiesta := auto{marca: "Ford", color: "Azul", anio: 2021, modelo: "turbo"}
	fmt.Println(fiesta)

	var corsa auto
	corsa.marca = "Cherolet"
	fmt.Println(corsa)

}
*/

func main() {
	var fiesta auto.AutoPublico
	fiesta.Marca = "Ford"
	fmt.Println(fiesta)
}
