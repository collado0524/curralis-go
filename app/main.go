package main

import (
	"fmt"
	"os"
)

func main() {
	// Verificar si se han pasado argumentos
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <argumento1> <argumento2> ...")
		return
	}

	// Imprimir todos los argumentos
	fmt.Println("Argumentos pasados al programa:")
	for i, arg := range os.Args {
		fmt.Printf("Argumento %d: %s\n", i, arg)
	}

	// Obtener el primer argumento (excluyendo el nombre del programa)
	if len(os.Args) > 1 {
		arg1 := os.Args[1]
		fmt.Println("Primer argumento:", arg1)
	}

	// Puedes usar os.Args[1:], para obtener todos los argumentos menos el nombre del programa
	args := os.Args[1:]
	fmt.Println("Argumentos sin el nombre del programa:", args)
}
