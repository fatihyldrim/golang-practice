/*
	Go dili içinde C çalıştırma
*/
package main

//#include <stdio.h>
//void callC() {
//    printf("This is from C code!\n");
//}
import "C"

import "fmt"

func main() {
	fmt.Println("A")
	C.callC()
	fmt.Println("B")
}
