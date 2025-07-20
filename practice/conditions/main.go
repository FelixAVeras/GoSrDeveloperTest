package main

import "fmt"

func main() {
	var fruta string

	var allowedAge int = 32
	var permission bool = false

	//if else statement
	if allowedAge >= 65 {
		fmt.Println("Es una persona en edad de jubilacion")
	} else if allowedAge >= 21 && permission {
		fmt.Println("Adelante, es mayor de edad y tiene permiso")
	} else if allowedAge >= 21 && !permission {
		fmt.Println("Adelante, es mayor de edad, pero no tiene permiso")
	} else {
		fmt.Println("Es menor de edad, no puede entrar")
	}

	//Switch statement
	fmt.Print("Ingrese el nombre de una fruta: ")
	fmt.Scan(&fruta)

	switch fruta {
	case "Manzana":
		fmt.Println("Ingresaste manzana")
	case "Pera":
		fmt.Println("Ingresaste pera")
	case "guineo":
		fmt.Println("Ingresaste banana")
	default:
		fmt.Println("No reconozco ese valor")
	}
}
