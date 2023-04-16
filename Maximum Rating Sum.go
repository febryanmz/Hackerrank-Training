package main

import "fmt"

const (
	UintSize = 32 << (^uint(0) >> 32 & 1)
	MaxInt   = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	MinInt   = -MaxInt - 1
)

func main() {
	input := []int{4, 5, -3}
	output := maxSubArray(input)
	fmt.Println(output)

	input = []int{1, 2, -4, 4, 1}
	output = maxSubArray(input)
	fmt.Println(output)
}

func maxSubArray(arr []int) int {
	lenArr := len(arr)
	var currentMax int
	max := MinInt
	for i := 0; i < lenArr; i++ {
		currentMax = currentMax + arr[i]
		if currentMax > max {
			max = currentMax
		}

		if currentMax < 0 {
			currentMax = 0
		}

	}
	return max
}
