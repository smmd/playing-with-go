// Parametros 1 2 3 4
package main

import (
	"fmt"
	"os"
)

// Recibir parámetros de consola e imprimirlos de regreso

func main() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	fmt.Println("Ejemplo de 'slice' cortar el elemento 0")
	// [primer elemento : ultimo elemento]
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}

	fmt.Println("Ejemplo de 'slice' limitado")
	// [primer elemento : ultimo elemento]
	for _, arg := range os.Args[1:3] {
		fmt.Println(arg)
	}
}
