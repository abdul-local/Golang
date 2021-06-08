package main

import "fmt"

func main() {

	//pengisian nilai array
	// var keluarga [4]string
	// keluarga[0] = "abdul"
	// keluarga[1] = "rodi"
	// keluarga[2] = "hamzan"
	// keluarga[3] = "huswatun"

	// fmt.Println(keluarga[0],keluarga[1],keluarga[2],keluarga[3])

	var fruits = [4]string{"apple", "orange", "banana", "mango"}
	fmt.Printf("isi semua element: %s\n", fruits)
	fmt.Printf("jumlah semua element: %d\n", len(fruits))

}
