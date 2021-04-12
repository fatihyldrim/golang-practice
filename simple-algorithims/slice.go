/*
	Slice
	Fonksiyonlara clice göndermek dizi göndermekten daha iyi dir.
	Slice ların otomatik adresleri gider ama dizilerin kopyası oluşturulur.
	Diziler genişleyemez.

	tanımlama yaparken eleman sayısı belirtirsen o bir dizi olur. Belirtmezse slcie olarak devam eder.

	copy, cap ve len

	array slice dönüşümü için array[0:] , array[:] kullanılabilir.
	copy fonsiyonu sadece slice alır.

	slice nil'e eşitlenirse len ve cap sıfırlanır. Tekrardan make gerekir.

*/
package main

import (
	"fmt"
	"sort"
)

func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func main() {
	s1 := []int{-1, 0, 2, 5}

	fmt.Print("s1: ")
	printSlice(s1)

	fmt.Printf("Cap: %d, Length: %d\n", cap(s1), len(s1))
	s1 = append(s1, -100)

	fmt.Printf("s1: ")
	printSlice(s1)
	fmt.Printf("Cap: %d, Length: %d\n", cap(s1), len(s1))

	s1 = append(s1, -2)
	s1 = append(s1, -3)
	s1 = append(s1, -4)
	printSlice(s1)
	fmt.Printf("Cap: %d, Length: %d\n", cap(s1), len(s1))

	s2 := [5]int{5, -5, 5, -5, 5}
	s3 := []int{7, 7, 7, -7, 7}
	copy(s2[0:], s3)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)

	s4 := make([]int, 3)
	fmt.Println(s4)

	s4 = nil
	fmt.Println(s4)

	//sort
	s5 := []int{1, 5, 3, 0, 12, 38, 18, 2}
	fmt.Println("-: ", s5)
	sort.Slice(s5, func(i, j int) bool {
		return s5[i] > s5[j]
	})
	fmt.Println(">: ", s5)

}
