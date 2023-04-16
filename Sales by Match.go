package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
    // Write your code here
	var sockMap = make(map[int32]int32, 0)
	var pair int32 = 0
	for _, v := range ar {
		sockMap[v]++
		if v := sockMap[v]; v % 2 == 0 {
			pair++
		}
	}
	return pair

}

func main()  {
	fmt.Println(sockMerchant(8, []int32{1,1,2,3,2,1,2,2,4}))
}