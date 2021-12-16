package main

import "fmt"

func main() {
	edad := 24
	empleado := true
	antiguedad := 2
	var sueldo float32 = 170000

	if(puedePedirPrestamo(edad, empleado, antiguedad)) {
		if(tieneInteres(sueldo)) {
			fmt.Printf("Puede pedir prestamo con interes.")
		} else {
			fmt.Printf("Puede pedir prestamo sin interes.")
		}
	} else {
		fmt.Printf("No puede pedir prestamo.")
	}
}

func puedePedirPrestamo(edad int, empleado bool, antiguedad int) bool {
	if(edad > 22 && empleado && antiguedad > 1) {
		return true
	} else {
		return false
	}
}

func tieneInteres(sueldo float32) bool {
	if(sueldo <= 100000) {
		return true
	}
	return false
}