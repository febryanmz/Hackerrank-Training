package main

import (
	"fmt"
	"log"
)

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var count int32 = 0
	if n < 2 {
		log.Fatal("n terlalu kecil")
	}
	for i := 0; i < len(ar)-1; i++ {
		for j := 1; j < len(ar); j++ {
			if i < j && (ar[i]+ar[j])%k == 0 {
				fmt.Println("Hasil = ", ar[i], " + ", ar[j])
				count++
			}
		}

	}
	return count

}

func main() {
	ar := []int32{1, 3, 2, 6, 1, 2}
	fmt.Println(divisibleSumPairs(6, 3, ar))
	ar2 := []int32{5, 9, 10, 7, 4}
	fmt.Println(divisibleSumPairs(5, 2, ar2))
}
