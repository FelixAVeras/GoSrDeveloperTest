package main

import "fmt"

func main() {
	myMaps := map[string]string{
		"Dominican Republic": "Santo Domingo",
		"Colombia":           "Bogota",
		"Puerto Rico":        "San Juan",
	}

	fmt.Println(myMaps)

	fmt.Println("La capital de Republica Dominicana es: ", myMaps["Dominican Republic"])

	myMaps["Brasil"] = "Brasilia"

	fmt.Println(myMaps)
}
