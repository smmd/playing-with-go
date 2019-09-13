// Parametros 1 2 3 4
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cadena := strings.Join(os.Args[1:], ",")
	fmt.Println(cadena)
}
