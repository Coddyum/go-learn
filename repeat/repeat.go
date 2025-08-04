package main

import "fmt"

func Repeat(l string, t int) (string) {
	result := ""
	for i := 0; i < t; i++ {
		result = result + l
	}
	return result
}

func main() {
	fmt.Println(Repeat("", 2))
}