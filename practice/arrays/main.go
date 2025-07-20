package main

import "fmt"

func main() {
	//Listado de frutas
	var fruitList = [4]string{
		"Manzana",
		"Pera",
		"Naranja",
		"Banana",
	}

	fmt.Println(fruitList)

	//Slice
	var countryList = []string{
		"Dominican Republic",
		"Mexico",
		"Brasil",
		"Colombia",
	}

	countryList = append(countryList, "United States")

	fmt.Println(countryList)

	//Rango en listas
	var countryList2 = countryList[1:3]
	fmt.Println(countryList2)

	var countryList3 = countryList[1:]
	fmt.Println(countryList3)

	var countryList4 = countryList[:3]
	fmt.Println(countryList4)
}
