/*
	runes

	Tüm stringler aslında bir byte slice larıdır. Her bir karakter bir byte

	struct larda eklenebilir.

	x hex, boşluk x boşluk hex,
*/
package main

import "fmt"

func main() {

	const s1 = "go123GO"
	fmt.Println("s1: ", s1)
	fmt.Printf("q: %q\n", s1)
	fmt.Printf("+q: %+q\n", s1)
	fmt.Printf("x: %x\n", s1)
	fmt.Printf(" x: % x\n", s1)
	fmt.Printf("s: %s\n", s1)
	fmt.Printf("v: %v\n", s1)
	fmt.Printf("#v: %#v\n", s1)
	fmt.Printf("d: %d\n", len(s1))

	const s2 = "\x67\x6f\x31\x32\x33\x47\x4f"
	fmt.Println("s2: ", s2)
	fmt.Printf("q: %q\n", s2)
	fmt.Printf("+q: %+q\n", s2)
	fmt.Printf("x: %x\n", s2)
	fmt.Printf(" x: % x\n", s2)
	fmt.Printf("s: %s\n", s2)
	fmt.Printf("v: %v\n", s2)
	fmt.Printf("#v: %#v\n", s2)
	fmt.Printf("d: %d\n", len(s2))
}
