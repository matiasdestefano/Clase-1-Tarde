package main

import "fmt"

func main() {
	var palabra string = "Programacion"
	fmt.Println("La palabra",palabra,"tiene una cantidad de:", len(palabra), "letras")

	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%c", palabra[i])
	}
}