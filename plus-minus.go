package main

import "fmt"

func cuk(arr []int32) float64 { // parameter
	// Write your code here
	var result float64
	var bilPos, bilNeg, bilNol int
	var TotalBil = len(arr) // total bilangan yang ada di array

	for i := 0; i < TotalBil; i++ {
		if arr[i] > 0 {
			bilPos++
		} else if arr[i] < 0 {
			bilNeg++
		} else {
			bilNol++
		}
	}

	PrintPos := float64(bilPos) / float64(TotalBil)
	PrintNeg := float64(bilNeg) / float64(TotalBil)
	PrintNol := float64(bilNol) / float64(TotalBil)

	fmt.Printf("%.6f\n", PrintPos) // print format 6 bilangan dibelakang NOL
	fmt.Printf("%.6f\n", PrintNeg)
	fmt.Printf("%.6f\n", PrintNol)

	// fmt.Printf("%.6f\n %.6f\n %.6f\n", positifRatio, negatifRatio, zeroRatio)

	return result

}

func main() {
	array := []int32{6, 7, 9, 11, 12, -7, -30, 0, -2, -3}
	fmt.Println("coba cuk, ojo bingung : ", cuk(array))
}
