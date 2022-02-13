package main

import (
	"fmt"
	"strings"
)

func validarPalidromo(texto string) string {
	var textoAlreves string
	fmt.Printf("Texto %s: %d\n", texto, len(texto))
	for i := (len(texto)) - 1; i >= 0; i-- {
		textoAlreves += string(texto[i])
	}

	if strings.ToUpper(textoAlreves) == strings.ToUpper(texto) {
		return "Palidromo"
	} else {
		return "No Palidromo"
	}
}

func main() {

	fmt.Println(validarPalidromo("Amor a Roma"))

}
