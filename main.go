package main

import "fmt"

func main() {
	jumlahSisi := 3
	for i := 1; i <= jumlahSisi; i++ {
		// cetak luarnya
		for j := jumlahSisi; j > i; j-- {
			fmt.Print(" ")
		}

		// cetak dalamnya
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("1")
		}
		fmt.Println()
	}
}

////
//    5
//

//5
//4
//3
