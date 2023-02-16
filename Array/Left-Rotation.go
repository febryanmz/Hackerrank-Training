package main

import "fmt"

func rotateLeft(d int, arr []int32) []int32 {
	var rst []int32
	aMin := 1
	aMax := 100000
	aTotal := len(arr)
	if aTotal < aMin || aTotal > aMax || int(d) < aMin || int(d) > aTotal {
		return rst
	}
	if aTotal == int(d) {
		return arr
	}
	fmt.Println("d :", d)
	iMax := 1000000
	left := arr[d:]
	fmt.Println("left :", left)
	right := arr[:d]
	fmt.Println("right :", right)
	for _, i := range left {
		if i < 1 || i > int32(iMax) {
			return nil
		}
		rst = append(rst, i)
	}
	for _, i := range right {
		if i < 1 || i > int32(iMax) {
			return nil
		}
		rst = append(rst, i)
	}

	return rst

}

func main() {
	fmt.Println(rotateLeft(2, []int32{1, 2, 3, 4, 5}))
	fmt.Println(rotateLeft(4, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
