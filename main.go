package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	aSecret := os.Getenv("A_SECRET")
	fmt.Println(aSecret)
	fmt.Println(reverseString(aSecret))
}

func reverseString(s string) string {
	r := []rune(s)
	sb := strings.Builder{}
	for i := len(r) - 1; i >= 0; i-- {
		sb.WriteRune(r[i])
	}
	return sb.String()
}
