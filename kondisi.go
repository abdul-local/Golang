package main

import "fmt"

func main() {

	var point = 10

	if point == 10 {
		fmt.Println("Nilai Sempurna")
	} else if point > 7 {
		fmt.Println("Nilai Lulus")
	} else {
		fmt.Println("Anda tidak lulus ")
	}

}
