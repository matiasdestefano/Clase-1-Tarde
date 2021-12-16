package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	var nombre string = "Benjamin"

	fmt.Printf("Edad de Benjamin: %d\n", employees[nombre])
	i := 0
	for _, edad := range employees {
		if(edad > 21) {
			i++
		}
	}
	fmt.Printf("La cantidad de empleados con edad mayor a 21 es de: %d", i)

	//Agregamos a Federico
	employees["Federico"] = 25

	//Eliminamos a Pedro
	delete(employees, "Pedro")
}