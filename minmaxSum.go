package main

import (
	"fmt"
	"math"
)

func miniMaxSum(arr []int32) {
	// var array []int // array tampungan
	// var i, j int    // variabel sementara

	// panjangarray := int(len(arr)) // panjang dari array (index ->angka)

	// for i = 0; i < panjangarray; i++ {
	// 	var sum int32
	// 	for j = 0; j < panjangarray; j++ {
	// 		if i != j {
	// 			sum += arr[j]
	// 		}
	// 	}
	// 	array = append(array, int(sum))
	// }

	// sort.Ints(array)
	// fmt.Println(array[0], array[panjangarray-1])

	//--------------
	sum := int64(0)
	min := int64(math.MaxInt64)
	max := int64(0)

	for _, v := range arr {
		sum += int64(v)

		if int64(v) < min {
			min = int64(v)
		}
		if int64(v) > max {
			max = int64(v)
		}
	}
	fmt.Printf("%d %d\n", sum-max, sum-min) // ini
}

func main() {
	miniMaxSum([]int32{1, 3, 5, 7, 9})          // 16 24
	miniMaxSum([]int32{2, 7, 8, 5, 4, 9, 3, 1}) // 30 38
	miniMaxSum([]int32{9, 10, 7, 6, 444, 87})   // 119 557
}
