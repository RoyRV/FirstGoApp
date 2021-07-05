package main

import (
	"fmt"
)

func main() {
	//A map (or dictionary) is an unordered collection of key-value pairs, where each key is unique.
	//The default zero value of a map is nil. A nil map is equivalent to an empty map except that elements can’t be added.
	arreglo := map[string]int{"key": 13}
	fmt.Println(arreglo)        // el arreglo
	fmt.Println(arreglo["key"]) //se puede acceder al valor a traves de la llave

	arreglo["key"] = 27 //se puede asignar un valor a la llave con los corchetes
	fmt.Println(arreglo["key"])

	//para borrar un eleento  puede utilizar el metodo delete
	delete(arreglo, "key")
	fmt.Println(arreglo)

	arreglo["key2"] = 9
	fmt.Println(arreglo)

	//se puede verificar si un valor existe o no
	if value, found := arreglo["key33"]; found {
		fmt.Println(value)
	} else {
		fmt.Println("valor no encontrado")
	}

	if arreglo["key33"] == 0 {
		fmt.Println("Objecto  NO en el arreglo")
	} else {
		fmt.Println("objeto en el arreglo")
	}
	arreglo["key3"] = 33
	arreglo["key4"] = 44
	arreglo["key5"] = 55

	// Como hacer un for each sobre un arreglo
	for key, value := range arreglo {
		fmt.Println(key, value)
	}
}
