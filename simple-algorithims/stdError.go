/*
Standart error çıktısı verir. Normal output ile farkını UNIX bash ayırır.
bash'de sadece hataları görmek için:

>go run stdErr.go > /tmp/stdError

*/

package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Missing argument"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
