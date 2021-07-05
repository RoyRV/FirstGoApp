package main

import "fmt"

//arreglos
func main() {
	var array [3]int
	array[0] = 1
	array[1] = 3
	array[2] = 2

	fmt.Println(array[1])

	//declaracion implicita
	array2 := [3]int{11, 12, 13}
	fmt.Println(array2)

	//slices
	slice1 := array2[:]
	fmt.Println(slice1)

	slice2 := []int{2, 4, 6, 7}
	fmt.Println((slice2))

	slice2 = append(slice2, 9)
	fmt.Println(slice2)

	//crear un slice (va a apuntar a los valores en el slice2)
	slice3 := slice2[1:]
	fmt.Println(slice3)
	slice4 := slice2[:2]
	fmt.Println(slice4)
	slice5 := slice2[2:8]
	fmt.Println(slice5)
}
