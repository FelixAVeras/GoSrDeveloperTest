package main

import "fmt"

func main() {
	var suma int = 0

	for i := 0; i <= 1000; i++ {
		if i%2 != 0 {
			suma = suma + i
		}
	}

	fmt.Println("La suma de los primeros 100 numeros es: ", suma)
}
