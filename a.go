package main

import "fmt"

func main() {
	var S string
	var hasil_penjumlahan, A, B int

	fmt.Scanln(&S, &A, &B)
	hasil_penjumlahan = A + B

	fmt.Printf("Kata =, %s\n", S)
	fmt.Printf("Jumlah = %d\n", hasil_penjumlahan)
}