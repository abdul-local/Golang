// package main

// import "fmt"

// func main() {

// perulangan menggunakan kyword for

// for i := 0; i <= 10; i++ {
// 	fmt.Printf("hasil: %d\n", i)
// }

// var i = 0

// for i < 10 {
// 	fmt.Printf("hasil :%d\n", i)
// 	i += 1

// }

// for {
// 	fmt.Printf("nilai :%d\n", i)
// 	i++
// 	if i == 5 {
// 		break
// 	}
// }

// for i := 0; i <= 10; i++ {
// 	if i%2 == 1 {
// 		continue
// 	}
// 	if i > 8 {
// 		break
// 	}

// 	fmt.Printf("nilai:%d\n", i)

// }

// perulangan bersarang
// for i := 0; i <= 10; i++ {

// 	for j := i; j < 10; j++ {

// 		fmt.Print(j, " ")
// 	}
// 	fmt.Println(i)
// }

// outerLoop:
// 	for i := 0; i < 5; i++ {
// 		for j := 0; j < 5; j++ {
// 			if j == 3 {
// 				break outerLoop
// 			}

// 			fmt.Print("matriks [", i, "][", j, "]", "\n")

// 		}
// 	}

// }
