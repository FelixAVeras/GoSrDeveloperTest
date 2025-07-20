package main

import (
	"fmt"
	"strconv"
)

func main() {
	//variable de tipo string
	var name string = "Felix"
	var middleName = "Antonio"
	lastname := "Carvajal"

	fmt.Println("Hola " + name + ", estas trabajando con Go!!!")

	//Variables de tipo numericas
	var currentYear int16 = 2025
	var shortYear int8 = 24
	myAge := 32

	fmt.Println("El año actual es: ", currentYear)
	fmt.Println("El año abreviado es: ", shortYear)
	fmt.Println("Mi edad es: ", myAge)

	fmt.Println("Mi nombre es " + name + " " + middleName + " " + lastname + " y mi edad es " + strconv.Itoa(myAge))
}
