package main

import "fmt"

func main() {
	s := "привет"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))
}
