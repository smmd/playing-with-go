package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "https://google.com"

	//Si usamos _ en lugar de err ignora el error que es lo que devuelve Get en segundo lugar
	r, err := http.Get(url)

	//Manejo de errores
	if err != nil {
		fmt.Printf("Error visitando %s, Ay %v", url, err)
		return
	}

	fmt.Println(r)

	//Liberar memoria
	defer r.Body.Close()

	body, err2 := ioutil.ReadAll(r.Body)

	if err2 != nil {
		fmt.Printf("Error %s", err)
		return
	}

	//fmt.Println(body) <- bits
	fmt.Printf("%s", body)

	for _, arg := range os.Args[1:] {
		response, _ := http.Get(arg)
		//body2, _ := ioutil.ReadAll(response.Body)

		fmt.Printf("%v", response)
	}
}
