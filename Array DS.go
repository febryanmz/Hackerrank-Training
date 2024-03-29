package main

import "fmt"

func reverseArray(a []int32) []int32 {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a

}

func main() {
	fmt.Println(reverseArray([]int32{1, 2, 3, 4}))
}
