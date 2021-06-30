package main

import "fmt"

/*
puedo ejecutar este codigo de dos maneras
-  go run FirstGoApp (el nombre del modulo)
-  go run *NombreArchivo*.go
*/
func main() {
	//declaracion tipo 1
	var texto string
	texto = "gaaa1"
	fmt.Println(texto)
	//declaracion tipo 2
	var texto2 string = "gaaaaaa2"
	fmt.Println(texto2)
	//declaracion tipo 3
	texto3 := "gaaaaa3"
	fmt.Println(texto3)
}
