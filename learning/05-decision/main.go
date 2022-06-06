package main

import "log"

func main() {
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"
	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 100
	if myNum > 99 && isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum == 100 || isTrue {
		log.Println("1")
	}

	myVar := "cat"
	switch myVar {
	case "cat":
		log.Println("Cat")
	case "dog":
		log.Println("Dog")
	default:
		log.Println("Other")
	}

}
