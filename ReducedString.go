package main

import "fmt"


func getMinDeletions(s string) int32 {
	// Write your code here
	var hitung int 
	var langkah int
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1]
		hitung += 1		
	} else {
		langkah += (hitung / 2)
		hitung += 1
	}
	langkah += (hitung/2)
	return langkah

}

func main()  {
	string s = "jakartaraya";
	fmt.Println(getMinDeletions)
}