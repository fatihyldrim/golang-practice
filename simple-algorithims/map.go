/*
	İki çeşit map tanımlama
	Map içinde key var mı sorgusu
	Map den item silme
	For ile map içinde dönme

	new ile make arasındaki fark: new poinder döner
*/
package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["k1"] = 1
	m1["k2"] = 2
	fmt.Println("m1:", m1)
	delete(m1, "k1")
	fmt.Println("m1:", m1)

	m2 := map[string]int{
		"k1": 1,
		"k2": 2,
	}
	fmt.Println("m2:", m2)

	//Key var mı?
	_, ok := m1["notInMap"]
	if ok {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Does NOT exist")
	}

	for key, value := range m2 {
		fmt.Println(key, value)
	}
}
