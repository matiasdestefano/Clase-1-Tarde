package main

import "fmt"

func main() {
	var precio float32 = 1000
	var descuento float32 = 20

	resultado := calcularDescuento(precio, descuento)
	fmt.Printf("El precio total es de: $%.2f", resultado)
}

func calcularDescuento(precio float32, descuento float32) float32 {
	return precio - ((precio * descuento) / 100)
}