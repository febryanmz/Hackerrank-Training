package main

import (
	"fmt"
	"strings"
)

var dnaComplement = strings.NewReplacer(
	"A", "T", "T", "A", "G", "C", "C", "G",
	"a", "t", "t", "a", "g", "c", "c", "g",
)

// reverse complementary DNA sequence
func rcDNA(s string) string {
	c := dnaComplement.Replace(s)
	rc := make([]byte, len(c))
	for i, j := 0, len(rc)-1; i < len(rc); i, j = i+1, j-1 {
		rc[i] = c[j]
	}
	return string(rc)
}

func main() {
	// TgggcAT from ATgcccA
	fmt.Println(rcDNA("ATgcccA"))
}
