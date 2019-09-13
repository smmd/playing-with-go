// Parametros , 1 2 3 4
package main

import (
	"fmt"
	"os"
)

// Recibir parámetros de consola e imprimirlos de regreso

func main() {
	// Leer argumentos del programa
	// Imprimir argumentos del programa

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	fmt.Println("Hasta aqui los parametros")

	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%v ", os.Args[i])
	}

	fmt.Println("Hasta aqui separados por un espacio")

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	fmt.Println("Hasta aqui con linea nueva")

	sep := os.Args[1]
	for i := 2; i < len(os.Args); i++ {
		fmt.Printf("%v "+sep, os.Args[i])
	}

	var imprimir string
	for i := 2; i < len(os.Args); i++ {
		imprimir += os.Args[i] + sep
	}

	fmt.Println(imprimir)
}
