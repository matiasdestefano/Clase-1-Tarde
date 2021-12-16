package main

import "fmt"

func main() {
	//Se puede hacer con Maps, Array o con un Switch
	var numero int = 1
	meses := map[int]string {1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio",
	7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre",}

	fmt.Printf("%s", meses[numero])
}