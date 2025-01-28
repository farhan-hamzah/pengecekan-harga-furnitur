package main

import "fmt"

func main() {
	var masukan int
	var keluaran float64

	fmt.Scan(&masukan)

	if masukan > 1000 {
		keluaran = float64(masukan) * 0.8
		fmt.Println(keluaran)
	} else if masukan >= 500 && masukan <= 1000 {
		keluaran = float64(masukan) * 0.95
		fmt.Println(keluaran)
	}
}