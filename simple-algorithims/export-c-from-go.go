/*
Go da yazılan fonksiyonları C kütüphanesi olacak build etmek.
//export fonksiyon adı yazılmak zorundadır.
>go build -o test.o -buildmode=c-shared test.go
*/
package main

import "C"

import (
	"fmt"
)

//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function!")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

func main() {
}
