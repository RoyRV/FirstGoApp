// ejemplo dos : POINTERS

package main

import "fmt"

func main() {
	//declarar un pointer , se tiene q inicializar
	//para asignarle un valor al pointer , tiene q utilizar
	//el asterisco antes del nombre de la variable
	var firstName *string = new(string)
	*firstName = "Test"
	fmt.Println(firstName, *firstName)

	//otro ejemplo,
	//creo una variable de tipo sting
	secondName := "Prueba2"
	fmt.Println(secondName)
	//creo una variable pointer apuntando a una variable string
	pointerSecondName := &secondName
	fmt.Println(pointerSecondName, secondName)
	//actualizo el valor de la variable
	//el valor del pointer no cambia
	secondName = "Prueba2-actualizado"
	fmt.Println(pointerSecondName, secondName)

	//constantes
	const pi = 3.1416
	fmt.Println(pi)
	//constante con tipo de dato
	const edad int = 14
	fmt.Println(edad + 2)
	//esto va a fallar
	//fmt.Println(edad + 3.13)

}
