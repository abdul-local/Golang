// package main

// import "fmt"

// func main() {

// var point = 10

// if point == 10 {
// 	fmt.Println("Nilai Sempurna")
// } else if point > 7 {
// 	fmt.Println("Nilai Lulus")
// } else {
// 	fmt.Println("Anda tidak lulus ")
// }

// var point = 8840.0

// if persent := point / 100; persent >= 100 {
// 	fmt.Printf("%.1f%s perfect\n", persent, "%")

// } else if persent >= 70 {
// 	fmt.Printf("%.1f%s good\n", persent, "%")

// } else {
// 	fmt.Printf("%.1f%s not bad\n", persent, "%")

// }

// seleksi menggunakan switch
// var nilai = 8
// switch nilai {
// case 10:
// 	fmt.Println("perfect")
// case 8:
// 	fmt.Println("good")

// default:
// 	fmt.Println("not bad")

// }

// pemanfaatan case utk banyak kondisi
// var nilai = 6
// switch nilai {
// case 10, 9:
// 	fmt.Println("perfect")
// case 8, 7:
// 	fmt.Println("good")
// case 6, 5:
// 	fmt.Println("not bad")

// default:
// 	fmt.Println("very bad")

// }

// swith dengan gaya if-else
// var point = 6

// switch {
// case point == 8:
// 	fmt.Println("perfect")
// case point <= 8 || point > 6:
// 	fmt.Println("Good")
// default:
// 	{
// 		fmt.Println("not bad")
// 		fmt.Println("You need to learn more")
// 	}
// }

// penggunaan fallthrough

// 	var point = 6

// 	switch {
// 	case point == 8:
// 		fmt.Println("perfect")
// 	case point <= 8 || point > 6:
// 		fmt.Println("Good")
// 		fallthrough
// 	case point < 5:
// 		fmt.Println("YOU NEEF TO LEARN MORE")
// 	default:
// 		{
// 			fmt.Println("not bad")
// 			fmt.Println("You need to learn more")
// 		}
// 	}

// }

// seleksi kondisi bersarang

// var point = 10
// if point > 9 {
// 	switch point {
// 	case 10:
// 		fmt.Println("perfect")
// 	default:
// 		fmt.Println("good")

// 	}

// } else {
// 	if point == 5 {
// 		fmt.Print("NOT BAD")

// 	} else if point == 3 {
// 		fmt.Println("keep trying")

// 	} else {
// 		fmt.Println("You can do it")
// 		if point == 0 {
// 			fmt.Println("You must to Hard learn from now")
// 		}

// 	}

// }

// }
