package main

import "fmt"

func main() {
	jumlahSisi := 10
	//jumlahSisi *= 2
	//jumlahSisi += 6

	jumlahSisi *= 2
	for i := jumlahSisi; i >= 1; i-- {
		// pengecekan dari tengah
		//for l := 0; l < l; l++ {
		//	if jumlahSisi/i == 0 {
		//		fmt.Println("*")
		//	}
		//}
		start := i / 2
		sisa := i % 2
		start += sisa
		if sisa != 1 {
			continue
		}
		for l := 1; l <= jumlahSisi/2; l++ {
			if start == l {
				fmt.Print("1")
			}

			if l < start {
				fmt.Print(" ")
			}

			//if start == l {
			//	fmt.Print("*")
			//} else {
			//	fmt.Print(" ")
			//}
		}
		for l := jumlahSisi / 2; l >= 1; l-- {
			//if start != l {
			//	fmt.Print(" ")
			//} else {
			//	fmt.Print("*")
			//}

			if start != l {
				fmt.Print(" ")
			}

			if l > start {
				fmt.Print("1")
			}
		}

		fmt.Println("")
	}

	//dataBagi2 := 9 / 2
	//dataSisa := 9 % dataBagi2
	//
	//fmt.Println(dataSisa)
}

////
//    5
//

//5
//4
//3
