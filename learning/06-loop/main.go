package main

import "log"

func main() {
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	animals := make(map[string]string)
	animals["dog"] = "Fido"
	animals["cat"] = "Fluffy"
	for animalType, animal := range animals {
		log.Println(animalType, animal)
	}

	var firstLine = "Once upon a midnight dreary"
	for i, l := range firstLine {
		log.Println(i, ":", l)
	}
}
