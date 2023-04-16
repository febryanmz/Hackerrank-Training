package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	// bikin variabel penampung
	var sum1 int32
	var sum2 int32
	var hasil int32
	// baris berisi (n) bilangan bulat
	// n=5 (5 * 5) bidang
	// arr[i][j] // [i]=kolom // [j]==baris
	for i := 0; i < len(arr); i++ {
		sum1 += arr[i][i]
		// fmt.Println(sum1)
		sum2 += arr[i][(len(arr)-1)-i] /// n-1 harusnya hmmm
		// fmt.Println(sum2)
	}
	hasil = int32(math.Abs(float64(sum1 - sum2)))
	{
		return hasil
	}

	// baris ke [0] 11, 2, 4
	// baris ke [1] 4, 5, 6
	// baris ke [2] 10, 8, -12
	//------
	// kolom [0] 11, 4, 10
	// kolom [1] 2, 5, 8
	// kolom [2] 4, 6, -12
	// 0 arr[0][0], arr[0][1], arr[0][2]
	// 1 arr[1][0], arr[1][1], arr[1][2]
	// 2 arr[2][0], arr[2][1], arr[2][2]
	//------
	// jumlah arr[0][0], arr[1][1], arr[2][2] <<< ini
	// jumlah arr[0][2], arr[1][1], arr[2][0]
}
func main() {
	fmt.Println(diagonalDifference([][]int32{{6, 3, 9}, {8, 10, 15}}))
}
